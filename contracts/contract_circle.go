package contracts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/scriptnull/badgeit/common"
)

// CircleBadgeContract checks for circle ci badges
type CircleBadgeContract struct {
	Path string
}

// NewCircleBadgeContract returns contract for checking various badges related to circle ci
func NewCircleBadgeContract(path string) *CircleBadgeContract {
	return &CircleBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for circle ci
func (contract CircleBadgeContract) Badges() ([]common.Badge, error) {
	// check if circle.yml exists
	_, circleYmlErr := os.Stat(filepath.Join(contract.Path, "circle.yml"))

	// check if .circleci folder exists
	_, circleFolderErr := os.Stat(filepath.Join(contract.Path, ".circleci"))

	if os.IsNotExist(circleYmlErr) && os.IsNotExist(circleFolderErr) {
		return []common.Badge{}, nil
	}

	repos := common.GetGithubRepos(contract.Path)
	shieldsURL := "https://img.shields.io/circleci/project"
	circleURL := "https://circleci.com"

	var badges []common.Badge

	for _, repo := range repos {
		chatRoom := &common.Badge{
			Name:     "circle badge",
			ImageURL: fmt.Sprintf("%s/github/%s.svg", shieldsURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/gh/%s", circleURL, repo.Slug),
			Group:    "circle ci",
			Tags:     []string{"circle", "circleci", "continuous integration", "ci", "testing", "tests", "build", "status"},
		}
		badges = append(badges, *chatRoom)
	}

	return badges, nil
}
