package main

import (
	"flag"
	"fmt"
	"os"

	version "github.com/gitXite/pulse/loadtester/internal/version"
)

func main() {
	// global option flags
	versionFlag := flag.Bool("version", false, "Print version information and exit.")
	flag.BoolVar(versionFlag, "v", false, "Alias for version information.")

	helpFlag := flag.Bool("help", false, "Print help menu and exit.")
	flag.BoolVar(helpFlag, "h", false, "Alias for help menu")

	// run command and option flags
	if len(os.Args) > 1 && os.Args[1] == "run" {
		runFlags := flag.NewFlagSet("run", flag.ExitOnError)

		url := runFlags.String("url", "http://localhost:8080", "Target URL to test.")
		runFlags.StringVar(url, "u", "http://localhost:8080", "Alias for target URL.")

		requests := runFlags.Int("requests", 100, "Total number of requests.")
		runFlags.IntVar(requests, "r", 100, "Alias for total requests.")

		concurrency := runFlags.Int("concurrency", 10, "Number of concurrent workers.")
		runFlags.IntVar(concurrency, "c", 10, "Alias for concurrent workers.")

		runFlags.Parse(os.Args[2:])
		fmt.Printf("Running pulse on %s with %d requests and %d concurrent workers\n", *url, *requests, *concurrency)
		return
	}

	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Version)
		return
	}
	if *helpFlag {
		fmt.Printf("Pulse %s\n", version.Version)
		fmt.Println("A lightweight load-testing toolkit")
		fmt.Println()
		fmt.Println("USAGE:\n	pulse [GLOBAL OPTIONS] <COMMAND> [COMMAND OPTIONS] [ARGUMENTS]")
		fmt.Println()
		fmt.Println("COMMANDS:")
		fmt.Println("	run   Execute the core worker processes.")
		fmt.Println()
		fmt.Println("COMMAND OPPTIONS:")
		fmt.Println("	-u, --url			Target URL to test.")
		fmt.Println("	-r, --requests		Total number of requests.")
		fmt.Println("	-c, --concurrency	Number of concurrent workers.")
		// fmt.Println("	-C, --config		Read from a .yaml config file.")
		fmt.Println()
		fmt.Println("GLOBAL OPTIONS:")
		fmt.Println("	-h, --help     Print this help menu and exit.")
		fmt.Println("	-v, --version  Print version information and exit.")
		// fmt.Println("	-q, --quiet    Suppress all standard output logs.")
		return
	}

}
