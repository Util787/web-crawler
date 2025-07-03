package main

import (
	"log/slog"
	"os"

	"github.com/Util787/task-manager/pkg/logger/handlers/slogpretty"
	v2 "github.com/Util787/web-crawler/internal/app/v2"
)

func main() {
	v2.Run(setupPrettySlog())
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
