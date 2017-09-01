package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

// CodeclimateBadgeContract checks for Codeclimate badges
type CodeclimateBadgeContract struct {
	Path string
}

// NewCodeclimateBadgeContract returns contract for checking various badges related to Codeclimate
func NewCodeclimateBadgeContract(path string) *CodeclimateBadgeContract {
	return &CodeclimateBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for Codeclimate
func (contract CodeclimateBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)
	codeclimateBadgesURL := "https://img.shields.io/codeclimate/github"
	codeclimateCoverageURL := "https://img.shields.io/codeclimate/coverage/github"
	codeclimateIssuesURL := "https://img.shields.io/codeclimate/issues/github"

	codeclimateURL := "https://codeclimate.com/github"

	var badges []common.Badge

	// version
	for _, repo := range repos {
		versionBadge := &common.Badge{
			Name:     "Code Climate",
			ImageURL: fmt.Sprintf("%s/%s.svg", codeclimateBadgesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", codeclimateURL, repo.Slug),
			Tags:     []string{"version"},
		}

		badges = append(badges, *versionBadge)
	}

	// coverage
	for _, repo := range repos {
		coverageBadge := &common.Badge{
			Name:     "Code Climate",
			ImageURL: fmt.Sprintf("%s/%s.svg", codeclimateCoverageURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s/coverage", codeclimateURL, repo.Slug),
			Tags:     []string{"coverage", "percent"},
		}

		badges = append(badges, *coverageBadge)
	}

	// issues
	for _, repo := range repos {
		issuesBadge := &common.Badge{
			Name:     "Code Climate",
			ImageURL: fmt.Sprintf("%s/%s.svg", codeclimateIssuesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s/issues", codeclimateURL, repo.Slug),
			Tags:     []string{"issues"},
		}

		badges = append(badges, *issuesBadge)
	}
	commonTags := []string{"codeclimate", "code coverage", "static analysis", "coverage"}
	for i := 0; i < len(badges); i++ {
		badges[i].Group = "codeclimate"
		badges[i].Tags = append(badges[i].Tags, commonTags...)
	}

	return badges, nil
}
