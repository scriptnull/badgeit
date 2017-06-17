package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

const (
	shieldBaseURL = "https://img.shields.io/github/downloads"
)

type GithubBadgeContract struct {
	Path string
}

func NewGithubBadgeContract(path string) *GithubBadgeContract {
	return &GithubBadgeContract{
		Path: path,
	}
}

func (contract GithubBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)

	fmt.Printf("%v", repos)
	return []common.Badge{}, nil
}

func generateDownloadBadges(repo common.GithubRepo) (badges []common.Badge) {
	allReleases := &common.Badge{
		Name: "github all releases",
	}
	allReleases.Markdown = fmt.Sprintf("[![%s](%s/%s/total.svg)](%s)", allReleases.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *allReleases)

	latestRelease := &common.Badge{
		Name: "github latest release",
	}
	latestRelease.Markdown = fmt.Sprintf("[![%s](%s/%s/latest/total.svg)](%s)", latestRelease.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *latestRelease)

	return
}

func generateVersionBadges(repo common.GithubRepo) (badges []common.Badge) {
	return
}

func generateSocialBadges(repo common.GithubRepo) (badges []common.Badge) {
	return
}

func generateMiscBadges(repo common.GithubRepo) (badges []common.Badge) {
	return
}
