package commands

// command names
const (
	HelpCommand        = "help"
	CrawlCommand       = "crawl"
	ToFileCommand      = "tofile"
	ToTerminalCommand  = "toterm"
	ShowParamsCommand  = "show"
	ResetParamsCommand = "reset"
)

// command usages
var (
	cmdUsages = []string{
		CrawlPageUsage,
		ToFileUsage,
		ToTerminalUsage,
		ShowParamsUsage,
		ResetParamsUsage,
	}
)
