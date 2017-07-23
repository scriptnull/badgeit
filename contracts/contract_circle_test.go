package contracts

import (
	"testing"
)

func TestCircleBadges(t *testing.T) {
	var testMap = map[string]bool{
		"[![circle badge](https://img.shields.io/circleci/project/github/circleci/frontend.svg)](https://circleci.com/gh/circleci/frontend)": false,
	}
	contract := NewCircleBadgeContract("../samples/circle")
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
