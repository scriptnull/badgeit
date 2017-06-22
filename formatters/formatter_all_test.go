package formatters

import (
	"testing"

	"github.com/scriptnull/badgeit/common"
)

func TestFormat(t *testing.T) {
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
	allFormatter := AllFormatter{
		FormatterOption{
			Badges: badges,
		},
	}
	result := allFormatter.Format()
	expected := "[![sample badge 1](imageurl 1)](linkurl 1) [![sample badge 2](imageurl 2)](linkurl 2)"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Basic case with badge and delimiter
	allFormatter.Delimiter = " xx "
	result = allFormatter.Format()
	expected = "[![sample badge 1](imageurl 1)](linkurl 1) xx [![sample badge 2](imageurl 2)](linkurl 2)"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
