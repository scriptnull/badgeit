package contracts

import (
	"testing"
)

func TestTravisBadges(t *testing.T) {
	var testMap = map[string]bool{
		"[![travis badge](https://img.shields.io/travis/rust-lang/cargo.svg)](https://travis-ci.org/rust-lang/cargo)": false,
	}
	contract := NewTravisBadgeContract("../samples/travis")
	badges, err := contract.Badges()
	if err != nil {
		t.Error("Expected err to be not nil, but got", err)
	}

	for _, badge := range badges {
		testMap[badge.Markdown()] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}
