package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scriptnull/badgeit/formatters"
)

func main() {
	// Parse Flags
	fFlag := flag.String("-f", "all", "Format for arranging the badges.")
	_ = flag.String("-d", " ", "Delimiter to be used.")
	flag.Parse()

	// Get Suitable Formatter
	formatter, err := formatters.NewFormatter(*fFlag)
	if err != nil {
		fmt.Fprint(os.Stderr, "Invalid -f option.")
	}

	// Print result
	result := formatter.Format()
	fmt.Println(result)
	os.Exit(0)
}
