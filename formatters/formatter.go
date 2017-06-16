package formatters

import "errors"
import "github.com/scriptnull/badgeit/common"

// Formatter formats the badges in desired fashion
type Formatter interface {
	Format() string
}

// FormatterOption has options for formatter
type FormatterOption struct {
	CmdArgType string
	Badges     []common.Badge
	Delimiter  string
}

// NewFormatter chooses the required formatter
func NewFormatter(formatterArg FormatterOption) (Formatter, error) {
	switch formatterArg.CmdArgType {
	case "all":
		return AllFormatter{formatterArg}, nil
	case "min":
		return MinFormatter{formatterArg}, nil
	}
	return AllFormatter{formatterArg}, errors.New("Unknown formatter")
}
