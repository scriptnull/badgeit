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
