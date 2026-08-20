package main

import (
	"fmt"

	version "github.com/gitXite/pulse/loadtester/internal/version"
)

func printHelp() {
	fmt.Printf("Pulse %s\n", version.Version)
	fmt.Println("A lightweight load-testing toolkit")
	fmt.Println()

	fmt.Println("USAGE:\n	pulse [GLOBAL OPTIONS] <COMMAND> [COMMAND OPTIONS]")
	fmt.Println()

	fmt.Println("COMMANDS:")
	fmt.Println("	run              Execute the core worker processes.")
	fmt.Println("	start            Start the TUI for Pulse.")
	fmt.Println()

	fmt.Println("COMMAND OPPTIONS:")
	fmt.Println("	-u, --url        Target URL to test.")
	fmt.Println("	-d, --duration   How long the test runs.")
	fmt.Println("	-w, --workers    Number of concurrent workers.")
	// fmt.Println("	-r, --requests		Total request limit.")
	// fmt.Println("	-R, --rate			Set throughput target.")
	// fmt.Println("	-m, --method		Set HTTP method.")
	// fmt.Println("	-c, --config		Read from a .yaml config file.")
	// fmt.Println("	-f, --format		Select output format.")
	// fmt.Println("	-o, --output		Set output destination.")
	fmt.Println()

	fmt.Println("GLOBAL OPTIONS:")
	fmt.Println("	-h, --help       Print this help menu and exit.")
	fmt.Println("	-v, --version    Print version information and exit.")
	// fmt.Println("	-q, --quiet    Suppress all standard output logs.")
}
