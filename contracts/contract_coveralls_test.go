package contracts

import (
	"testing"
)

func TestGenerateCoverallsBadge(t *testing.T) {
	var testMap = map[string]bool{
		"[![Coveralls](https://img.shields.io/coveralls/sindresorhus/xo.svg)](https://coveralls.io/github/sindresorhus/xo)": false,
	}
	contract := NewCoverallsBadgeContract("../samples/coveralls/valid")
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

func TestInvalidCoverallsBadge(t *testing.T) {
	contract := NewCoverallsBadgeContract("../samples/coveralls/invalid")
	badges, err := contract.Badges()
	if err != nil {
		t.Error("Expected err to be not nil, but got", err)
	}

	if len(badges) == 0 {
		t.Errorf("Expected to generate badges for coveralls disabled repos as well")
	}
}
