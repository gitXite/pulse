package main

import (
	"flag"
	"fmt"

	version "github.com/gitXite/pulse/loadtester/internal/version"
)

func main() {
	// command option flags
	url := flag.String("url", "http://localhost:8080", "Target URL to test.")
	flag.StringVar(url, "u", "http://localhost:8080", "Alias for target URL.")
	requests := flag.Int("requests", 100, "Total number of requests.")
	flag.IntVar(requests, "r", 100, "Alias for total requests.")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent workers.")
	flag.IntVar(concurrency, "c", 10, "Alias for concurrent workers.")

	// global option flags
	versionFlag := flag.Bool("version", false, "Print version information and exit.")
	flag.BoolVar(versionFlag, "v", false, "Alias for version information.")
	
	flag.Parse()

	if (*versionFlag) {
		fmt.Println(version.Version)
		return
	}

	fmt.Printf("Running pulse on %s with %d requests and %d concurrent workers\n", *url, *requests, *concurrency)
	
	// fmt.Printf("Pulse v%s\n", "dev")
	// fmt.Println("A lightweight load-testing toolkit")
	// fmt.Println()
	// fmt.Println("USAGE:\n    pulse [GLOBAL OPTIONS] <COMMAND> [COMMAND OPTIONS] [ARGUMENTS]")
	// fmt.Println()
	// fmt.Println("COMMANDS:")
	// fmt.Println("    run   Execute the core worker processes.")
	// fmt.Println()
	// fmt.Println("COMMAND OPPTIONS:")
	// fmt.Println("    -u, --url			Target URL to test.")
	// fmt.Println("    -r, --requests		Total number of requests.")
	// fmt.Println("    -c, --concurrency	Number of concurrent workers.")
	// fmt.Println()
	// fmt.Println("GLOBAL OPTIONS:")
	// fmt.Println("    -h, --help     Print this help menu and exit.")
	// fmt.Println("    -v, --version  Print version information and exit.")
	// fmt.Println("    -q, --quiet    Suppress all standard output logs.")
	
}