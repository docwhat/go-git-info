package main

import (
	"fmt"

	"docwhat.org/go-git-info/internal/app/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}

func main() {
	cli.Run(versionString())
}
