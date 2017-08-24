package common

import (
	"fmt"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// GithubRepo : Contains the parsed github URL details
type GithubRepo struct {
	Username string
	RepoName string
	Slug     string
	URL      string
}

// GetGithubRepos : Function to parse the github repo
func GetGithubRepos(path string) (githubRepos []GithubRepo) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return
	}
	httpsBase, sshBase := "https://github.com/", "git@github.com:"
	for _, remote := range remotes {
		err = remote.Config().Validate()
		if err == nil {
			url := remote.Config().URL
			trimmedURL := strings.TrimSuffix(url, ".git")

			var repoSlug string
			if strings.Contains(trimmedURL, httpsBase) {
				repoSlug = strings.Replace(trimmedURL, httpsBase, "", 1)
			}
			if strings.Contains(trimmedURL, sshBase) {
				repoSlug = strings.Replace(trimmedURL, sshBase, "", 1)
			}
			splittedSlug := strings.Split(repoSlug, "/")
			if len(splittedSlug) == 2 {
				githubRepo := GithubRepo{
					Slug:     repoSlug,
					Username: splittedSlug[0],
					RepoName: splittedSlug[1],
					URL:      fmt.Sprintf("%s%s", httpsBase, repoSlug),
				}
				githubRepos = append(githubRepos, githubRepo)
			}
		}
	}
	return
}
