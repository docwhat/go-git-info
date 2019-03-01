package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}

func main() {
	fmt.Printf("version: %v\n", versionString())
}
