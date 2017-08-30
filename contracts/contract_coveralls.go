package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

// CoverallsBadgeContract checks for Coveralls badges
type CoverallsBadgeContract struct {
	Path string
}

// NewCoverallsBadgeContract returns contract for checking various badges related to Coveralls
func NewCoverallsBadgeContract(path string) *CoverallsBadgeContract {
	return &CoverallsBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for Coveralls
func (contract CoverallsBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)
	coverallsBadgesURL := "https://img.shields.io/coveralls" // 
	coverallsURL := "https://coveralls.io/github"

	var badges []common.Badge
	for _, repo := range repos {
		masterBadge := &common.Badge{
			Name:     "Coveralls",
			ImageURL: fmt.Sprintf("%s/%s.svg", coverallsBadgesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", coverallsURL, repo.Slug),
			Group:    "coveralls",
			Tags:     []string{"coveralls", "code coverage", "static analysis", "coverage"},
		}

		badges = append(badges, *masterBadge)
	}

	return badges, nil
}
