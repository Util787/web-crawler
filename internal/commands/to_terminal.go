package commands

import (
	"fmt"

	"github.com/Util787/web-crawler/internal/crawler"
)

var ToTerminalUsage = fmt.Sprintf("%v — write current pages to terminal", ToTerminalCommand)

func ToTerminal(c *crawler.Crawler) {
	fmt.Println("Pages:")
	for page := range c.Pages {
		fmt.Println(page)
	}
}
