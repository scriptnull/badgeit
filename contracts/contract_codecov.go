package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

// CodecovBadgeContract checks for Codecov badges
type CodecovBadgeContract struct {
	Path string
}

// NewCodecovBadgeContract returns contract for checking various badges related to Codecov
func NewCodecovBadgeContract(path string) *CodecovBadgeContract {
	return &CodecovBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for Codecov
func (contract CodecovBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)
	codecovBadgesURL := "https://img.shields.io/codecov/c/github"
	codecovURL := "https://codecov.io/gh"

	var badges []common.Badge
	for _, repo := range repos {
		masterBadge := &common.Badge{
			Name:     "Codecov",
			ImageURL: fmt.Sprintf("%s/%s.svg", codecovBadgesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", codecovURL, repo.Slug),
			Group:    "codecov",
			Tags:     []string{"codecov", "code coverage", "static analysis", "coverage"},
		}

		badges = append(badges, *masterBadge)
	}

	return badges, nil
}
