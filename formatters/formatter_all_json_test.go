package formatters

import (
	"testing"

	"github.com/scriptnull/badgeit/common"
)

func TestAllJSONFormat(t *testing.T) {
	badges := []common.Badge{
		common.Badge{
			Name:     "sample badge 1",
			ImageURL: "imageurl 1",
			LinkURL:  "linkurl 1",
		},
		common.Badge{
			Name:     "sample badge 2",
			ImageURL: "imageurl 2",
			LinkURL:  "linkurl 2",
		},
	}

	// Basic case with only badges
	f := AllJSONFormatter{
		FormatterOption{
			Badges: badges,
		},
	}
	result := f.Format()
	expected := `[{"Name":"sample badge 1","ImageURL":"imageurl 1","LinkURL":"linkurl 1","Style":"","Label":"","MarkdownContent":"[![sample badge 1](imageurl 1)](linkurl 1)"},{"Name":"sample badge 2","ImageURL":"imageurl 2","LinkURL":"linkurl 2","Style":"","Label":"","MarkdownContent":"[![sample badge 2](imageurl 2)](linkurl 2)"}]`
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
