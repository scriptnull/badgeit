package contracts

import (
	"testing"
)

func TestGenerateGitterBadge(t *testing.T) {
	var testMap = map[string]bool{
		"[![gitter chat room](https://badges.gitter.im/scriptnull/badgeit.svg)](https://gitter.im/scriptnull/badgeit)": false,
	}
	contract := NewGitterBadgeContract("../samples/gitter")
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
