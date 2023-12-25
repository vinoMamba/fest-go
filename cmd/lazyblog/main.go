package main

import (
	"os"

	"github.com/vinoMamba.com/lazyblog/internal/lazyblog"
)

func main() {
	command := lazyblog.NewLazyBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
