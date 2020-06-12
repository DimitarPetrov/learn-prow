package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(strings.ToLower(str))
}

func parseArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("usage: lower [string]")
	}
	if len(args) != 1 {
		return "", fmt.Errorf("expected exactly one argument, found %d", len(args))
	}
	return args[0], nil
}
