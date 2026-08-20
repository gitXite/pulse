package main

import (
	"flag"
	"fmt"
	"os"
)

// runs pulse cli
func run() {
	runFlags := flag.NewFlagSet("run", flag.ExitOnError)

	url := runFlags.String("url", "http://localhost:8080", "target URL to test.")
	runFlags.StringVar(url, "u", "http://localhost:8080", "alias for target URL.")

	duration := runFlags.String("duration", "60s", "how long the test runs.")
	runFlags.StringVar(duration, "d", "60s", "alias for duration.")

	workers := runFlags.Int("workers", 10, "number of concurrent workers.")
	runFlags.IntVar(workers, "w", 10, "alias for concurrent workers.")

	runFlags.Parse(os.Args[2:])
	fmt.Printf("Running pulse on %s for %s, with %d concurrent workers\n", *url, *duration, *workers)
}
