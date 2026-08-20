package version

import "fmt"

var Version = "dev"

func PrintVersion() {
	fmt.Println(Version)
}
