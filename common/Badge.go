package common

import "fmt"

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
	imageURL := b.ImageURL
	if b.Style != "" {
		imageURL = fmt.Sprintf("%s?style=%s", imageURL, b.Style)
	}
	if b.Label != "" {
		imageURL = fmt.Sprintf("%s&label=%s", imageURL, b.Label)
	}
	return fmt.Sprintf("[![%s](%s)](%s)", b.Name, imageURL, b.LinkURL)
}
