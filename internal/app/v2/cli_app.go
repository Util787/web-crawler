package v2

import (
	"bufio"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Util787/web-crawler/internal/commands"
	"github.com/Util787/web-crawler/internal/common"
	"github.com/Util787/web-crawler/internal/crawler"
)

// TODO: test promptui on linux (it doesnt work good on windows)
const defaultOutputFilename = "output"

func Run(log *slog.Logger) {
	reader := bufio.NewReader(os.Stdin)

	var httpClientTimeout int
	var concurrencyLimit int
	var baseURL string
	var maxPages int

	fmt.Println("Use default params? (y/n): ")
	confirmInput, _ := reader.ReadString('\n')
	confirmInput = strings.ToLower(strings.TrimSpace(confirmInput))
	if strings.HasPrefix(confirmInput, "y") {
		httpClientTimeout = 10
		concurrencyLimit = 100
		maxPages = 10
		baseURL = "https://google.com"
	} else {
		httpClientTimeout, concurrencyLimit, baseURL, maxPages = commands.SetParams(reader)
	}

	c := crawler.New(time.Second*time.Duration(httpClientTimeout), log, baseURL, concurrencyLimit, maxPages)
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
				commands.ToFile(c, filename)
			} else {
				commands.ToTerminal(c)
			}

		case commands.ToFileCommand:
			var filename string

			if len(args) > 1 && args[1] != "" {
				filename = args[1]
			} else {
				fmt.Println("No filename provided. Using default:", defaultOutputFilename)
				filename = defaultOutputFilename
			}

			commands.ToFile(c, filename)

		case commands.ToTerminalCommand:
			commands.ToTerminal(c)

		case commands.ShowParamsCommand:
			commands.ShowParams(c, httpClientTimeout, concurrencyLimit, maxPages)

		case commands.ResetParamsCommand:
			httpClientTimeout, concurrencyLimit, baseURL, maxPages = commands.ResetParams(reader, &httpClientTimeout, &concurrencyLimit, &baseURL, &maxPages)
			c = crawler.New(time.Second*time.Duration(httpClientTimeout), log, baseURL, concurrencyLimit, maxPages)
			fmt.Println("Parameters updated and crawler re-initialized.")

		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

func getFilename(reader *bufio.Reader) string {
	fmt.Printf("Enter filename (skip with Enter to use default: '%s'): ", defaultOutputFilename)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("No filename provided. Using default:", defaultOutputFilename)
		return defaultOutputFilename
	}
	return input
}
