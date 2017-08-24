package contracts

import (
	"testing"
)

func TestGenerateCodecovBadge(t *testing.T) {
	var testMap = map[string]bool{
		"[![Codecov](https://img.shields.io/codecov/c/github/sindresorhus/make-dir.svg)](https://codecov.io/gh/sindresorhus/make-dir)": false,
	}
	contract := NewCodecovBadgeContract("../samples/codecov/valid")
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

func TestInvalidCodecovBadge(t *testing.T) {
	contract := NewCodecovBadgeContract("../samples/codecov/invalid")
	badges, err := contract.Badges()
	if err != nil {
		t.Error("Expected err to be not nil, but got", err)
	}

	if len(badges) == 0 {
		t.Errorf("Expected to generate badges for codecov disabled repos as well")
	}
}
