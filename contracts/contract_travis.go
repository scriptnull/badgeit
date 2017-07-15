package contracts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/scriptnull/badgeit/common"
)

// TravisBadgeContract checks for gitter badges
type TravisBadgeContract struct {
	Path string
}

// NewTravisBadgeContract returns contract for checking various badges related to gitter
func NewTravisBadgeContract(path string) *TravisBadgeContract {
	return &TravisBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for gitter
func (contract TravisBadgeContract) Badges() ([]common.Badge, error) {
	// check if .travis.yml exists
	if _, err := os.Stat(filepath.Join(contract.Path, ".travis.yml")); os.IsNotExist(err) {
		return []common.Badge{}, nil
	}

	repos := common.GetGithubRepos(contract.Path)
	shieldsURL := "https://img.shields.io/travis"
	travisURL := "https://travis-ci.org"

	var badges []common.Badge

	for _, repo := range repos {
		chatRoom := &common.Badge{
			Name:     "travis badge",
			ImageURL: fmt.Sprintf("%s/%s.svg", shieldsURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", travisURL, repo.Slug),
		}
		badges = append(badges, *chatRoom)
	}

	return badges, nil
}
