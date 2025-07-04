package commands

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Util787/web-crawler/internal/common"
)

var ResetParamsUsage = fmt.Sprintf("%v — change parameters (timeout, concurrency, baseURL) and remake crawler", ResetParamsCommand)

func SetParams(reader *bufio.Reader) (httpClientTimeout int, concurrencyLimit int, baseURL string) {
	httpClientTimeout = getHttpClientTimeout(reader)
	concurrencyLimit = getConcurrencyLimit(reader)
	baseURL = getBaseURL(reader)
	return httpClientTimeout, concurrencyLimit, baseURL
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
