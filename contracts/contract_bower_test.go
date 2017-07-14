package contracts

import "testing"

func TestBowerBadges(t *testing.T) {
	testMap := map[string]bool{
		"[![bower version](https://img.shields.io/bower/v/bootstrap.svg)](https://getbootstrap.com)": false,
		"[![license badge](https://img.shields.io/bower/l/bootstrap.svg)](https://getbootstrap.com)": false,
	}
	contract := NewBowerBadgeContract("../samples/bower")
	badges, err := contract.Badges()
	if err != nil {
		t.Fatal("Expected err to be nil, but got", err)
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
