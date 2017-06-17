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
	shieldBaseURL := "https://img.shields.io/github"

	tag := &common.Badge{
		Name: "github tag",
	}
	tag.Markdown = fmt.Sprintf("[![%s](%s/tag/%s.svg)](%s)", tag.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *tag)

	release := &common.Badge{
		Name: "github release",
	}
	release.Markdown = fmt.Sprintf("[![%s](%s/release/%s.svg)](%s)", release.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *release)

	preRelease := &common.Badge{
		Name: "github pre release",
	}
	preRelease.Markdown = fmt.Sprintf("[![%s](%s/release/%s/all.svg)](%s)", preRelease.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *preRelease)

	return
}

func generateSocialBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github"

	fork := &common.Badge{
		Name: "github fork",
	}
	fork.Markdown = fmt.Sprintf("[![%s](%s/forks/%s.svg?style=social&label=Fork)](%s)", fork.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *fork)

	stars := &common.Badge{
		Name: "github stars",
	}
	stars.Markdown = fmt.Sprintf("[![%s](%s/stars/%s.svg?style=social&label=Star)](%s)", stars.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *stars)

	watchers := &common.Badge{
		Name: "github watchers",
	}
	watchers.Markdown = fmt.Sprintf("[![%s](%s/watchers/%s.svg?style=social&label=Watch)](%s)", watchers.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *watchers)

	return
}

func generateMiscBadges(repo common.GithubRepo) (badges []common.Badge) {
	shieldBaseURL := "https://img.shields.io/github"

	openIssues := &common.Badge{
		Name: "github open issues",
	}
	openIssues.Markdown = fmt.Sprintf("[![%s](%s/issues/%s.svg)](%s)", openIssues.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *openIssues)

	closedIssues := &common.Badge{
		Name: "github closed issues",
	}
	closedIssues.Markdown = fmt.Sprintf("[![%s](%s/issues-closed/%s.svg)](%s)", closedIssues.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *closedIssues)

	openPR := &common.Badge{
		Name: "github open pr",
	}
	openPR.Markdown = fmt.Sprintf("[![%s](%s/issues-pr/%s.svg)](%s)", openPR.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *openPR)

	closedPR := &common.Badge{
		Name: "github closed pr",
	}
	closedPR.Markdown = fmt.Sprintf("[![%s](%s/issues-pr-closed/%s.svg)](%s)", closedPR.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *closedPR)

	contributors := &common.Badge{
		Name: "github contributors",
	}
	contributors.Markdown = fmt.Sprintf("[![%s](%s/contributors/%s.svg)](%s)", contributors.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *contributors)

	license := &common.Badge{
		Name: "github license",
	}
	license.Markdown = fmt.Sprintf("[![%s](%s/license/%s.svg)](%s)", license.Name, shieldBaseURL, repo.Slug, repo.URL)
	badges = append(badges, *license)

	return
}
