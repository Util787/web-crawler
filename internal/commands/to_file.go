package commands

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Util787/web-crawler/internal/crawler"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

var ToFileUsage = fmt.Sprintf("%v <filename> — write current pages to file", ToFileCommand)

func ToFile(c *crawler.Crawler, filename string) {
	// Check if file exists
	if _, err := os.Stat(filename + ".txt"); err == nil {
		// File exists, asking for action
		prompt := promptui.Select{
			Label: fmt.Sprintf("File %s already exists. Choose action:", filename),
			Items: []string{"Overwrite file", "Create new file", "Output to terminal"},
		}

		index, _, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				os.Exit(0)
			}
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
			ToTerminal(c)
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
