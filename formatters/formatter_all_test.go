package formatters

import (
	"testing"

	"github.com/scriptnull/badgeit/common"
)

func TestFormat(t *testing.T) {
	badges := []common.Badge{
		common.Badge{
			Name:     "sample badge 1",
			Markdown: "[describe1](link1)",
		},
		common.Badge{
			Name:     "sample badge 2",
			Markdown: "[describe2](link2)",
		},
	}

	// Basic case with only badges
	allFormatter := AllFormatter{
		FormatterOption{
			Badges: badges,
		},
	}
	result := allFormatter.Format()
	expected := "[describe1](link1) \n[describe2](link2)"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Basic case with badge and delimiter
	allFormatter.Delimiter = " xx "
	result = allFormatter.Format()
	expected = "[describe1](link1) xx [describe2](link2)"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
