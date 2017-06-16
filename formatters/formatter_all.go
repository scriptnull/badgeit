package formatters

import "strings"

// AllFormatter shows all the badges
type AllFormatter struct {
	FormatterOption
}

// Format gives a format in which all badges are present
func (fm AllFormatter) Format() string {
	// Set default Delimiter
	if fm.Delimiter == "" {
		fm.Delimiter = " \n"
	}

	// Generate badge markdown strings
	var badgeStr string
	for _, badge := range fm.Badges {
		badgeStr += badge.Markdown + fm.Delimiter
	}

	return strings.TrimSuffix(badgeStr, fm.Delimiter)
}
