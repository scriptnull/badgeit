package contracts

import "testing"

func TestBadges(t *testing.T) {
	testMap := map[string]bool{
		"[![npm weekly downloads](https://img.shields.io/npm/dw/express.svg)](https://npmjs.com/package/express)":         false,
		"[![npm monthly downloads](https://img.shields.io/npm/dm/express.svg)](https://npmjs.com/package/express)":        false,
		"[![npm yearly downloads](https://img.shields.io/npm/dy/express.svg)](https://npmjs.com/package/express)":         false,
		"[![npm total downloads](https://img.shields.io/npm/dt/express.svg)](https://npmjs.com/package/express)":          false,
		"[![npm version](https://img.shields.io/npm/v/express.svg)](https://npmjs.com/package/express)":                   false,
		"[![npm next version](https://img.shields.io/npm/v/express/next.svg)](https://npmjs.com/package/express)":         false,
		"[![npm canary version](https://img.shields.io/npm/v/express/canary.svg)](https://npmjs.com/package/express)":     false,
		"[![license badge](https://img.shields.io/npm/l/express.svg)](https://npmjs.com/package/express)":                 false,
		"[![snyk - known vulnerabilities](https://snyk.io/test/npm/express/badge.svg)](https://snyk.io/test/npm/express)": false,
	}
	contract := NewNpmBadgeContract("../samples/npm")
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
