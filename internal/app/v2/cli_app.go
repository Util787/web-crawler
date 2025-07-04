package v2

import (
	"bufio"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Util787/web-crawler/internal/commands"
	"github.com/Util787/web-crawler/internal/common"
	"github.com/Util787/web-crawler/internal/crawler"
)

// TODO: test promptui on linux (it doesnt work good on windows)
const defaultOutputFilename = "output.txt"

func Run(log *slog.Logger) {
	reader := bufio.NewReader(os.Stdin)

	var httpClientTimeout int
	var concurrencyLimit int

	httpClientTimeout = getHttpClientTimeout(reader)
	concurrencyLimit = getConcurrencyLimit(reader)

	c := crawler.New(time.Second*time.Duration(httpClientTimeout), log, getBaseURL(reader), concurrencyLimit)
	fmt.Println("Crawler initialized. Enter a command (type 'help' for a list of commands):")

	for {
		fmt.Print("> ")
		cmdLine, _ := reader.ReadString('\n')
		cmdLine = strings.TrimSpace(cmdLine)
		if cmdLine == "" {
			continue
		}
		args := strings.Fields(cmdLine)
		cmd := strings.ToLower(args[0])

		switch cmd {

		case commands.HelpCommand:
			commands.Help()

		case commands.CrawlCommand:
			// if url is provided, validate it and set as baseURL
			if len(args) > 1 {
				url := args[1]
				if err := common.ValidateURL(url); err != nil {
					fmt.Printf("Error:%v. Please enter a valid URL.\n", err)
					continue
				}
				if _, err := http.Get(url); err != nil {
					fmt.Printf("Error:%v. Can't ping url.\n", err)
					continue
				}
				c.BaseURL = url
			}
			// reset pages
			c.Pages = make(map[string]struct{})

			commands.CrawlPage(c, c.BaseURL)
			fmt.Printf("Crawling finished. Pages found: %d\n", len(c.Pages))

			fmt.Print("Make output to file? (y/n): ")
			confirmInput, _ := reader.ReadString('\n')
			confirmInput = strings.ToLower(strings.TrimSpace(confirmInput))

			if strings.HasPrefix(confirmInput, "y") {
				filename := getFilename(reader)
				makeOutputToFile(c, filename)
			} else {
				fmt.Println("Pages:")
				for page := range c.Pages {
					fmt.Println(page)
				}
			}

		case commands.ShowParamsCommand:
			fmt.Println("HTTP Client Timeout:", httpClientTimeout)
			fmt.Println("Concurrency Limit:", concurrencyLimit)
			fmt.Println("Base URL:", c.BaseURL)

		case commands.ResetParamsCommand:
			httpClientTimeout = getHttpClientTimeout(reader)
			concurrencyLimit = getConcurrencyLimit(reader)
			c = crawler.New(time.Second*time.Duration(httpClientTimeout), log, getBaseURL(reader), concurrencyLimit)
			fmt.Println("Parameters updated and crawler re-initialized.")

		case commands.ExitCommand:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

func getFilename(reader *bufio.Reader) string {
	fmt.Printf("Enter filename (skip with Enter to use default: %s): ", defaultOutputFilename)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultOutputFilename
	}
	return input
}

func makeOutputToFile(c *crawler.Crawler, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "Date:%s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(writer, "Base_URL:%s\n", c.BaseURL)
	for page := range c.Pages {
		fmt.Fprintf(writer, "%s\n", page)
	}
	writer.Flush()
	fmt.Printf("Output written to %s\n", filename)
}

func getHttpClientTimeout(reader *bufio.Reader) int {
	var res int
	for {
		fmt.Print("Enter http client timeout (int): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil || val <= 0 {
			fmt.Println("Error! Please enter a positive integer >0.")
			continue
		}
		res = val
		break
	}
	return res
}

func getConcurrencyLimit(reader *bufio.Reader) int {
	var res int
	for {
		fmt.Print("Enter concurrency limit (int): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil || val <= 0 {
			fmt.Println("Error! Please enter a positive integer >0.")
			continue
		}
		res = val
		break
	}
	return res
}

func getBaseURL(reader *bufio.Reader) string {
	var res string
	for {
		fmt.Print("Enter base URL: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Error! Please enter a non-empty URL.")
			continue
		}
		if err := common.ValidateURL(input); err != nil {
			fmt.Printf("Error:%v. Please enter a valid URL.\n", err)
			continue
		}
		if _, err := http.Get(input); err != nil {
			fmt.Printf("Error:%v. Can't ping url.\n", err)
			continue
		}
		res = input
		break
	}
	return res
}
