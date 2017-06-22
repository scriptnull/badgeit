package contracts

import (
	"fmt"

	"github.com/scriptnull/badgeit/common"
)

// GithubBadgeContract checks for github badges
type GithubBadgeContract struct {
	Path string
}

// NewGithubBadgeContract returns contract for checking various badges releated to github
func NewGithubBadgeContract(path string) *GithubBadgeContract {
	return &GithubBadgeContract{
		Path: path,
	}
}

// Badges returns the badges for github
func (contract GithubBadgeContract) Badges() ([]common.Badge, error) {
	repos := common.GetGithubRepos(contract.Path)

	var badges []common.Badge
	for _, repo := range repos {
		downloadBadges := generateDownloadBadges(repo)
		versionBadges := generateVersionBadges(repo)
		socialBadges := generateSocialBadges(repo)
		miscBadges := generateMiscBadges(repo)
		badges = append(badges, downloadBadges...)
		badges = append(badges, versionBadges...)
		badges = append(badges, socialBadges...)
		badges = append(badges, miscBadges...)
	}

	return badges, nil
}

func generateDownloadBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github/downloads"

	allReleases := &common.Badge{
		Name:     "github all releases",
		ImageURL: fmt.Sprintf("%s/%s/total.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *allReleases)

	latestRelease := &common.Badge{
		Name:     "github latest release",
		ImageURL: fmt.Sprintf("%s/%s/latest/total.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *latestRelease)

	return
}

func generateVersionBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github"

	tag := &common.Badge{
		Name:     "github tag",
		ImageURL: fmt.Sprintf("%s/tag/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *tag)

	release := &common.Badge{
		Name:     "github release",
		ImageURL: fmt.Sprintf("%s/release/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *release)

	preRelease := &common.Badge{
		Name:     "github pre release",
		ImageURL: fmt.Sprintf("%s/release/%s/all.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *preRelease)

	return
}

func generateSocialBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github"

	fork := &common.Badge{
		Name:     "github fork",
		ImageURL: fmt.Sprintf("%s/forks/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
		Style:    "social",
		Label:    "Fork",
	}
	badges = append(badges, *fork)

	stars := &common.Badge{
		Name:     "github stars",
		ImageURL: fmt.Sprintf("%s/stars/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
		Style:    "social",
		Label:    "Star",
	}
	badges = append(badges, *stars)

	watchers := &common.Badge{
		Name:     "github watchers",
		ImageURL: fmt.Sprintf("%s/watchers/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
		Style:    "social",
		Label:    "Watch",
	}
	badges = append(badges, *watchers)

	return
}

func generateMiscBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github"

	openIssues := &common.Badge{
		Name:     "github open issues",
		ImageURL: fmt.Sprintf("%s/issues/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *openIssues)

	closedIssues := &common.Badge{
		Name:     "github closed issues",
		ImageURL: fmt.Sprintf("%s/issues-closed/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *closedIssues)

	openPR := &common.Badge{
		Name:     "github open pr",
		ImageURL: fmt.Sprintf("%s/issues-pr/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *openPR)

	closedPR := &common.Badge{
		Name:     "github closed pr",
		ImageURL: fmt.Sprintf("%s/issues-pr-closed/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *closedPR)

	contributors := &common.Badge{
		Name:     "github contributors",
		ImageURL: fmt.Sprintf("%s/contributors/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *contributors)

	license := &common.Badge{
		Name:     "github license",
		ImageURL: fmt.Sprintf("%s/license/%s.svg", shieldBaseURL, repo.Slug),
		LinkURL:  repo.URL,
	}
	badges = append(badges, *license)

	return
}
