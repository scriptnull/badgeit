package contracts

import (
	"testing"

	"github.com/scriptnull/badgeit/common"
)

func TestGenerateDownloadBadges(t *testing.T) {
	repo := common.GithubRepo{
		Slug:     "atom/atom",
		Username: "atom",
		RepoName: "atom",
		URL:      "https://github.com/atom/atom",
	}
	testMap := map[string]bool{
		"[![github all releases](https://img.shields.io/github/downloads/atom/atom/total.svg)](https://github.com/atom/atom)":          false,
		"[![github latest release](https://img.shields.io/github/downloads/atom/atom/latest/total.svg)](https://github.com/atom/atom)": false,
	}

	badges := generateDownloadBadges(repo)
	for _, badge := range badges {
		testMap[badge.Markdown] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}

func TestGenerateVersionBadges(t *testing.T) {
	repo := common.GithubRepo{
		Slug:     "atom/atom",
		Username: "atom",
		RepoName: "atom",
		URL:      "https://github.com/atom/atom",
	}
	testMap := map[string]bool{
		"[![github tag](https://img.shields.io/github/tag/atom/atom.svg)](https://github.com/atom/atom)":                 false,
		"[![github release](https://img.shields.io/github/release/atom/atom.svg)](https://github.com/atom/atom)":         false,
		"[![github pre release](https://img.shields.io/github/release/atom/atom/all.svg)](https://github.com/atom/atom)": false,
	}

	badges := generateVersionBadges(repo)
	for _, badge := range badges {
		testMap[badge.Markdown] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}

func TestGenerateSocialBadges(t *testing.T) {
	repo := common.GithubRepo{
		Slug:     "atom/atom",
		Username: "atom",
		RepoName: "atom",
		URL:      "https://github.com/atom/atom",
	}
	testMap := map[string]bool{
		"[![github fork](https://img.shields.io/github/forks/atom/atom.svg?style=social&label=Fork)](https://github.com/atom/atom)":         false,
		"[![github stars](https://img.shields.io/github/stars/atom/atom.svg?style=social&label=Star)](https://github.com/atom/atom)":        false,
		"[![github watchers](https://img.shields.io/github/watchers/atom/atom.svg?style=social&label=Watch)](https://github.com/atom/atom)": false,
	}

	badges := generateSocialBadges(repo)
	for _, badge := range badges {
		testMap[badge.Markdown] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}

func TestGenerateMiscBadges(t *testing.T) {
	repo := common.GithubRepo{
		Slug:     "atom/atom",
		Username: "atom",
		RepoName: "atom",
		URL:      "https://github.com/atom/atom",
	}
	testMap := map[string]bool{
		"[![github open issues](https://img.shields.io/github/issues/atom/atom.svg)](https://github.com/atom/atom)":          false,
		"[![github closed issues](https://img.shields.io/github/issues-closed/atom/atom.svg)](https://github.com/atom/atom)": false,
		"[![github open pr](https://img.shields.io/github/issues-pr/atom/atom.svg)](https://github.com/atom/atom)":           false,
		"[![github closed pr](https://img.shields.io/github/issues-pr-closed/atom/atom.svg)](https://github.com/atom/atom)":  false,
		"[![github contributors](https://img.shields.io/github/contributors/atom/atom.svg)](https://github.com/atom/atom)":   false,
		"[![github license](https://img.shields.io/github/license/atom/atom.svg)](https://github.com/atom/atom)":             false,
	}

	badges := generateMiscBadges(repo)
	for _, badge := range badges {
		testMap[badge.Markdown] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}
