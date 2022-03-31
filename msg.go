package main

import (
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
)

func log(msg string) {
	fmt.Fprintf(os.Stdout, "%s: %s\n", green+"LOG"+reset, msg)
}

func warn(msg string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", yellow+"WARN"+reset, msg)
}

func exit(msg error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", red+"ERROR"+reset, msg)
	os.Exit(1)
}
