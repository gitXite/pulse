package main

import (
	"flag"
	"os"

	version "github.com/gitXite/pulse/loadtester/internal/version"
)

func main() {
	// global option flags
	versionFlag := flag.Bool("version", false, "print version information and exit.")
	flag.BoolVar(versionFlag, "v", false, "alias for version information.")

	helpFlag := flag.Bool("help", false, "print help menu and exit.")
	flag.BoolVar(helpFlag, "h", false, "alias for help menu")

	// check for commands
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			run() // parses flags and runs pulse
		case "start":
			start() // parses flags and starts pulse TUI
		}
	}

	flag.Parse()

	if *versionFlag {
		version.PrintVersion()
		return
	}
	if *helpFlag {
		printHelp()
		return
	}

}
