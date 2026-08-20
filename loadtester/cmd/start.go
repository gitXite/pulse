package main

import (
	"flag"
	"os"

	tui "github.com/gitXite/pulse/loadtester/internal/tui"
)

// starts pulse TUI
func start() {
	startFlags := flag.NewFlagSet("start", flag.ExitOnError)
	startFlags.Parse(os.Args[2:])

	tui.Start()
}
