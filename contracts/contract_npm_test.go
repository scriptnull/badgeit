package contracts

import "testing"

func TestBadges(t *testing.T) {
	testMap := map[string]bool{
		"[![npm weekly downloads](https://img.shields.io/npm/dw/vici.svg)](https://npmjs.com/package/vici)":         false,
		"[![npm monthly downloads](https://img.shields.io/npm/dm/vici.svg)](https://npmjs.com/package/vici)":        false,
		"[![npm yearly downloads](https://img.shields.io/npm/dy/vici.svg)](https://npmjs.com/package/vici)":         false,
		"[![npm total downloads](https://img.shields.io/npm/dy/vici.svg)](https://npmjs.com/package/vici)":          false,
		"[![npm version](https://img.shields.io/npm/v/vici.svg)](https://npmjs.com/package/vici)":                   false,
		"[![npm next version](https://img.shields.io/npm/v/vici/next.svg)](https://npmjs.com/package/vici)":         false,
		"[![npm canary version](https://img.shields.io/npm/v/vici/canary.svg)](https://npmjs.com/package/vici)":     false,
		"[![license badge](https://img.shields.io/npm/l/vici.svg)](https://npmjs.com/package/vici)":                 false,
		"[![snyk - known vulnerabilities](https://snyk.io/test/npm/vici/badge.svg)](https://snyk.io/test/npm/vici)": false,
	}
	contract := NewNpmBadgeContract("../samples/npm/basic")
	badges, err := contract.Badges()
	if err != nil {
		t.Fatal("Expected err to be nil, but got", err)
	}
	for _, badge := range badges {
		testMap[badge.Markdown] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}
