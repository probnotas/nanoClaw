package main

import (
	"fmt"
	"os"

	"github.com/probnotas/nanoClaw/cmd/nanoclaw-launcher-tui/internal/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
