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

// reset takes pointers to allow to keep current values if pointers are not nil
func ResetParams(reader *bufio.Reader, httpClientTimeoutPtr *int, concurrencyLimitPtr *int, baseURLPtr *string, maxPagesPtr *int) (httpClientTimeout int, concurrencyLimit int, baseURL string, maxPages int) {
	fmt.Println("Resetting parameters, press Enter to keep current value")
	httpClientTimeout = getHttpClientTimeout(reader, httpClientTimeoutPtr)
	concurrencyLimit = getConcurrencyLimit(reader, concurrencyLimitPtr)
	baseURL = getBaseURL(reader, baseURLPtr)
	maxPages = getMaxPages(reader, maxPagesPtr)
	return httpClientTimeout, concurrencyLimit, baseURL, maxPages
}

func getHttpClientTimeout(reader *bufio.Reader, httpClientTimeoutPtr *int) int {
	var res int
	for {
		fmt.Print("Enter http client timeout (int): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && httpClientTimeoutPtr != nil {
			res = *httpClientTimeoutPtr
			break
		}
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

func getConcurrencyLimit(reader *bufio.Reader, concurrencyLimitPtr *int) int {
	var res int
	for {
		fmt.Print("Enter concurrency limit (int): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && concurrencyLimitPtr != nil {
			res = *concurrencyLimitPtr
			break
		}
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

func getBaseURL(reader *bufio.Reader, baseURLPtr *string) string {
	var res string
	for {
		fmt.Print("Enter base URL: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			if baseURLPtr != nil {
				res = *baseURLPtr
				break
			}
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

func getMaxPages(reader *bufio.Reader, maxPagesPtr *int) int {
	var res int
	for {
		fmt.Print("Enter max pages (int, 0 for no limit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" && maxPagesPtr != nil {
			res = *maxPagesPtr
			break
		}
		val, err := strconv.Atoi(input)
		if err != nil || val < 0 {
			fmt.Println("Error! Please enter a positive integer >=0.")
			continue
		}
		res = val
		break
	}
	return res
}
