package formatters

// MinFormatter shows minimum number (upto 7) of essential badges
type MinFormatter struct {
	FormatterOption
}

// Format gives a format containing upto 7 essential badges
func (MinFormatter) Format() string {
	return "min formatter"
}
