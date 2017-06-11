package formatters

import "errors"

// Formatter formats the badges in desired fashion
type Formatter interface {
	Format() string
}

// NewFormatter chooses the required formatter
func NewFormatter(formatterArg string) (Formatter, error) {
	switch formatterArg {
	case "all":
		return AllFormatter{}, nil
	case "min":
		return MinFormatter{}, nil
	}
	return AllFormatter{}, errors.New("Unknown formatter")
}
