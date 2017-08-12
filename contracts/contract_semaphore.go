package contracts

import (
	"fmt"
	"net/http"

	"github.com/scriptnull/badgeit/common"
)

// SemaphoreBadgeContract checks for Semaphore badges
type SemaphoreBadgeContract struct {
	Path string
}

// NewSemaphoreBadgeContract returns contract for checking various badges related to Semaphore
func NewSemaphoreBadgeContract(path string) *SemaphoreBadgeContract {
	return &SemaphoreBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for Semaphore
func (contract SemaphoreBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)
	semaphoreBadgesURL := "https://semaphoreci.com/api/v1"
	semaphoreURL := "https://semaphoreci.com"

	var badges []common.Badge

	for _, repo := range repos {
		masterBadge := &common.Badge{
			Name:     "semaphore master badge",
			ImageURL: fmt.Sprintf("%s/%s/branches/master/badge.svg", semaphoreBadgesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", semaphoreURL, repo.Slug),
			Group:    "semaphore",
			Tags:     []string{"semaphore", "semaphoreci", "continuous integration", "ci", "testing", "tests", "build", "status"},
		}

		resp, err := http.Get(masterBadge.ImageURL)
		if err == nil && resp.StatusCode < 300 {
			badges = append(badges, *masterBadge)
		}
	}

	return badges, nil
}
