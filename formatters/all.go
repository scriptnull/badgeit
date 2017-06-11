package formatters

// AllFormatter shows all the badges
type AllFormatter struct{}

// Format gives a format in which all badges are present
func (AllFormatter) Format() string {
	return "all formatter"
}
