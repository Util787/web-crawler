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
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

// TODO: test promptui on linux (it doesnt work good on windows)
const defaultOutputFilename = "output"

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
				makeOutputToTerminal(c)
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
	// Check if file exists
	if _, err := os.Stat(filename + ".txt"); err == nil {
		// File exists, asking for action
		prompt := promptui.Select{
			Label: fmt.Sprintf("File %s already exists. Choose action:", filename),
			Items: []string{"Overwrite file", "Create new file", "Output to terminal"},
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Error choosing: %v\n", err)
			return
		}

		switch index {
		case 0: // Overwrite file
			writeToFile(c, filename)
		case 1: // Create new file
			newFilename := filename + "." + uuid.New().String()
			writeToFile(c, newFilename)
		case 2: // Output to terminal
			makeOutputToTerminal(c)
			return
		}
	} else {
		// File does not exist, creating new
		writeToFile(c, filename)
	}
}

func writeToFile(c *crawler.Crawler, filename string) {
	file, err := os.Create(filename + ".txt")
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
	fmt.Printf("Result written to %s.txt\n", filename)
}

func makeOutputToTerminal(c *crawler.Crawler) {
	fmt.Println("Pages:")
	for page := range c.Pages {
		fmt.Println(page)
	}
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
