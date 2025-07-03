package v2

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Util787/web-crawler/internal/commands"
	"github.com/Util787/web-crawler/internal/common"
	"github.com/Util787/web-crawler/internal/crawler"
)

func Run(log *slog.Logger) {
	reader := bufio.NewReader(os.Stdin)

	var httpClientTimeout int
	var concurrencyLimit int
	var baseURL string

	httpClientTimeout = getHttpClientTimeout(reader)
	concurrencyLimit = getConcurrencyLimit(reader)
	baseURL = getBaseURL(reader)

	c := crawler.New(time.Second*time.Duration(httpClientTimeout), log, baseURL, concurrencyLimit)
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
				baseURL = url
				c.BaseURL = baseURL
			}
			// reset pages
			c.Pages = make(map[string]struct{})

			commands.CrawlPage(c, c.BaseURL)
			fmt.Println("Pages:")
			for page := range c.Pages {
				fmt.Println(page)
			}
			fmt.Printf("Crawling finished. Pages found: %d\n", len(c.Pages))

		case commands.ShowParamsCommand:
			fmt.Println("HTTP Client Timeout:", httpClientTimeout)
			fmt.Println("Concurrency Limit:", concurrencyLimit)
			fmt.Println("Base URL:", c.BaseURL)

		case commands.ResetParamsCommand:
			httpClientTimeout = getHttpClientTimeout(reader)
			concurrencyLimit = getConcurrencyLimit(reader)
			baseURL = getBaseURL(reader)
			c = crawler.New(time.Second*time.Duration(httpClientTimeout), log, baseURL, concurrencyLimit)
			fmt.Println("Parameters updated and crawler re-initialized.")

		case commands.ExitCommand:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
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
		res = input
		break
	}
	return res
}
