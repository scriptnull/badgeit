package formatters

// AllFormatter shows all the badges
type AllFormatter struct {
	FormatterOption
}

// Format gives a format in which all badges are present
func (AllFormatter) Format() string {
	return "all formatter"
}
