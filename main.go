package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scriptnull/badgeit/common"
	"github.com/scriptnull/badgeit/contracts"
	"github.com/scriptnull/badgeit/formatters"
)

const VERSION = "0.1.0"

func main() {
	// Parse Flags
	fFlag := flag.String("f", "all", "Format for arranging the badges.")
	dFlag := flag.String("d", " ", "Delimiter to be used.")
	vFlag := flag.Bool("v", false, "Version information.")
	flag.Parse()

	if *vFlag {
		fmt.Println("badgeit version: " + VERSION)
		os.Exit(0)
	}
	// Obtain destination path, if not found, it will be cwd
	args := flag.Args()
	path, err := os.Getwd()
	if len(args) != 0 && len(args[0]) > 0 {
		path = args[0]
	}

	// Check Contract aggreement and obtain eligible badges
	var badges []common.Badge

	npmBadges, err := contracts.NewNpmBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, npmBadges...)
	}

	// Get Suitable Formatter
	formatter, err := formatters.NewFormatter(
		formatters.FormatterOption{
			Badges:    badges,
			Delimiter: *dFlag,
			Type:      *fFlag,
		},
	)
	if err != nil {
		fmt.Fprint(os.Stderr, "Invalid -f option.")
	}

	result := formatter.Format()
	fmt.Fprintf(os.Stdout, "%s\n", result)
	os.Exit(0)
}
