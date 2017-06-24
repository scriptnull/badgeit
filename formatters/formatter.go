package formatters

import "errors"
import "github.com/scriptnull/badgeit/common"

// Formatter formats the badges in desired fashion
type Formatter interface {
	Format() string
}

// FormatterOption has options for formatter
type FormatterOption struct {
	Type      string
	Badges    []common.Badge
	Delimiter string
	Style     string
}

// NewFormatter chooses the required formatter
func NewFormatter(formatterArg FormatterOption) (Formatter, error) {
	// Apply Style to be overidden
	if formatterArg.Style != "" {
		for i := 0; i < len(formatterArg.Badges); i++ {
			formatterArg.Badges[i].Style = formatterArg.Style
		}
	}

	// Choose formatter type
	switch formatterArg.Type {
	case "all":
		return AllFormatter{formatterArg}, nil
	case "min":
		return MinFormatter{formatterArg}, nil
	case "all-json":
		return AllJSONFormatter{formatterArg}, nil
	}
	return AllFormatter{formatterArg}, errors.New("Unknown formatter")
}
