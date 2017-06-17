package formatters

import (
	"fmt"
	"strings"
)

// AllFormatter shows all the badges
type AllFormatter struct {
	FormatterOption
}

// Format gives a format in which all badges are present
func (fm AllFormatter) Format() string {
	// Set default Delimiter
	if fm.Delimiter == "" {
		fm.Delimiter = " "
	}

	// replace escaped new line with actual new line
	fm.Delimiter = strings.Replace(fm.Delimiter, `\n`, "\n", -1)

	// Generate badge markdown strings
	var badgeStr string
	for _, badge := range fm.Badges {
		badgeStr = fmt.Sprintf("%s%s%s", badgeStr, badge.Markdown, fm.Delimiter)
	}

	return strings.TrimSuffix(badgeStr, fm.Delimiter)
}
