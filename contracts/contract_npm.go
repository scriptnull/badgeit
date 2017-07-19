package contracts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/scriptnull/badgeit/common"
)

// NpmBadgeContract checks for npm badges
type NpmBadgeContract struct {
	Path string
}

// NewNpmBadgeContract returns contract for checking various badges related to npm
func NewNpmBadgeContract(path string) *NpmBadgeContract {
	return &NpmBadgeContract{
		Path: path,
	}
}

// packageJSON is used for unmarshalling package.json file
type packageJSON struct {
	Name string
}

// Badges yields array of badges that satisfy the contract
func (contract *NpmBadgeContract) Badges() ([]common.Badge, error) {

	// read package.json
	raw, err := ioutil.ReadFile(filepath.Join(contract.Path, "package.json"))
	if err != nil {
		return []common.Badge{}, err
	}

	// parse json from package.json
	var pj packageJSON
	err = json.Unmarshal(raw, &pj)
	if err != nil {
		return []common.Badge{}, err
	}

	var badges []common.Badge

	if len(pj.Name) > 0 {
		shieldBaseURL := "https://img.shields.io/npm"
		npmBaseURL := "https://npmjs.com/package"

		// Download Badges

		weeklyDownloads := &common.Badge{
			Name:     "npm weekly downloads",
			ImageURL: fmt.Sprintf("%s/dw/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"weekly", "downloads", "count"},
		}
		badges = append(badges, *weeklyDownloads)

		monthlyDownloads := &common.Badge{
			Name:     "npm monthly downloads",
			ImageURL: fmt.Sprintf("%s/dm/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"monthly", "downloads", "count"},
		}
		badges = append(badges, *monthlyDownloads)

		yearlyDownloads := &common.Badge{
			Name:     "npm yearly downloads",
			ImageURL: fmt.Sprintf("%s/dy/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"yearly", "downloads", "count"},
		}
		badges = append(badges, *yearlyDownloads)

		totalDownloads := &common.Badge{
			Name:     "npm total downloads",
			ImageURL: fmt.Sprintf("%s/dt/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"total", "downloads", "count"},
		}
		badges = append(badges, *totalDownloads)

		// version Badges

		normalVersion := &common.Badge{
			Name:     "npm version",
			ImageURL: fmt.Sprintf("%s/v/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"version"},
		}
		badges = append(badges, *normalVersion)

		nextVersion := &common.Badge{
			Name:     "npm next version",
			ImageURL: fmt.Sprintf("%s/v/%s/next.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"next", "version"},
		}
		badges = append(badges, *nextVersion)

		canaryVersion := &common.Badge{
			Name:     "npm canary version",
			ImageURL: fmt.Sprintf("%s/v/%s/canary.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"canary", "version"},
		}
		badges = append(badges, *canaryVersion)

		// license badge
		licenseBadge := &common.Badge{
			Name:     "license badge",
			ImageURL: fmt.Sprintf("%s/l/%s.svg", shieldBaseURL, pj.Name),
			LinkURL:  fmt.Sprintf("%s/%s", npmBaseURL, pj.Name),
			Tags:     []string{"license"},
		}
		badges = append(badges, *licenseBadge)

		// snyk badge
		snykBadge := &common.Badge{
			Name:     "snyk - known vulnerabilities",
			ImageURL: fmt.Sprintf("https://snyk.io/test/npm/%s/badge.svg", pj.Name),
			LinkURL:  fmt.Sprintf("https://snyk.io/test/npm/%s", pj.Name),
			Tags:     []string{"snyk", "security", "vulnerability"},
		}
		badges = append(badges, *snykBadge)
	}

	commonTags := []string{"npm", "package manager"}
	for i := 0; i < len(badges); i++ {
		badges[i].Group = "npm"
		badges[i].Tags = append(badges[i].Tags, commonTags...)
	}

	return badges, nil
}
