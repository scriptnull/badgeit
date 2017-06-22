package common

import (
	"fmt"
	"strings"
)

// Badge is actual raw markdown badge
type Badge struct {
	Name     string
	ImageURL string
	LinkURL  string
	Style    string
	Label    string
}

// Markdown returns the markdown representation of the badge
func (b Badge) Markdown() string {
	var queryParams []string
	if b.Style != "" {
		queryParams = append(queryParams, "style="+b.Style)
	}
	if b.Label != "" {
		queryParams = append(queryParams, "label="+b.Label)
	}
	imageURL := b.ImageURL
	if len(queryParams) > 0 {
		imageURL += "?" + strings.Join(queryParams, "&")
	}
	return fmt.Sprintf("[![%s](%s)](%s)", b.Name, imageURL, b.LinkURL)
}
