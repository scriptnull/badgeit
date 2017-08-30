package contracts

import "testing"

func TestBadges(t *testing.T) {
	testMap := map[string]bool{
		"[![Code Climate](https://img.shields.io/codeclimate/issues/github/codeclimate/javascript-test-reporter.svg)](https://codeclimate.com/github/codeclimate/javascript-test-reporter/issues)":         false,
		"[![Code Climate](https://img.shields.io/codeclimate/github/codeclimate/javascript-test-reporter.svg)](https://codeclimate.com/github/codeclimate/javascript-test-reporter)":        false,
		"[![Code Climate](https://img.shields.io/codeclimate/coverage/github/codeclimate/javascript-test-reporter.svg)](https://codeclimate.com/github/codeclimate/javascript-test-reporter/coverage)":         false,	
	}
	contract := NewCodeclimateBadgeContract("../samples/codeclimate")
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
