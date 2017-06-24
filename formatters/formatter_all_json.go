package formatters

import "encoding/json"
import "github.com/scriptnull/badgeit/common"

// AllJSONFormatter shows all the badges in JSON format
type AllJSONFormatter struct {
	FormatterOption
}

type badgeJSON struct {
	common.Badge
	MarkdownContent string
}

// Format gives a format in which all badges are present
func (fm AllJSONFormatter) Format() string {
	var badges []badgeJSON

	for _, badge := range fm.Badges {
		b := badgeJSON{
			Badge:           badge,
			MarkdownContent: badge.Markdown(),
		}
		badges = append(badges, b)
	}

	result, err := json.Marshal(badges)
	if err != nil {
		return ""
	}

	return string(result)
}
