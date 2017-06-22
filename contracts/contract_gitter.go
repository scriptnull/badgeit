package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

// GitterBadgeContract checks for gitter badges
type GitterBadgeContract struct {
	Path string
}

// NewGitterBadgeContract returns contract for checking various badges related to gitter
func NewGitterBadgeContract(path string) *GitterBadgeContract {
	return &GitterBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for gitter
func (contract GitterBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)
	gitterBadgesURL := "https://badges.gitter.im"
	gitterURL := "https://gitter.im"

	var badges []common.Badge

	for _, repo := range repos {
		chatRoom := &common.Badge{
			Name:     "gitter chat room",
			ImageURL: fmt.Sprintf("%s/%s.svg", gitterBadgesURL, repo.Slug),
			LinkURL:  fmt.Sprintf("%s/%s", gitterURL, repo.Slug),
		}
		badges = append(badges, *chatRoom)
	}

	return badges, nil
}
