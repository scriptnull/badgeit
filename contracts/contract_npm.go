package contracts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/scriptnull/badgeit/common"
)

// NpmBadgeContract has the collection of contracts that yield one or more badges
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
			Name: "npm weekly downloads",
		}
		weeklyDownloads.Markdown = fmt.Sprintf("[![%s](%s/dw/%s.svg)](%s/%s)", weeklyDownloads.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *weeklyDownloads)

		monthlyDownloads := &common.Badge{
			Name: "npm monthly downloads",
		}
		monthlyDownloads.Markdown = fmt.Sprintf("[![%s](%s/dm/%s.svg)](%s/%s)", monthlyDownloads.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *monthlyDownloads)

		yearlyDownloads := &common.Badge{
			Name: "npm yearly downloads",
		}
		yearlyDownloads.Markdown = fmt.Sprintf("[![%s](%s/dy/%s.svg)](%s/%s)", yearlyDownloads.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *yearlyDownloads)

		totalDownloads := &common.Badge{
			Name: "npm total downloads",
		}
		totalDownloads.Markdown = fmt.Sprintf("[![%s](%s/dy/%s.svg)](%s/%s)", totalDownloads.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *totalDownloads)

		// version Badges

		normalVersion := &common.Badge{
			Name: "npm version",
		}
		normalVersion.Markdown = fmt.Sprintf("[![%s](%s/v/%s.svg)](%s/%s)", normalVersion.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *normalVersion)

		nextVersion := &common.Badge{
			Name: "npm next version",
		}
		nextVersion.Markdown = fmt.Sprintf("[![%s](%s/v/%s/next.svg)](%s/%s)", nextVersion.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *nextVersion)

		canaryVersion := &common.Badge{
			Name: "npm canary version",
		}
		canaryVersion.Markdown = fmt.Sprintf("[![%s](%s/v/%s/canary.svg)](%s/%s)", canaryVersion.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *canaryVersion)

		// license badge
		licenseBadge := &common.Badge{
			Name: "license badge",
		}
		licenseBadge.Markdown = fmt.Sprintf("[![%s](%s/l/%s.svg)](%s/%s)", licenseBadge.Name, shieldBaseURL, pj.Name, npmBaseURL, pj.Name)
		badges = append(badges, *licenseBadge)

		// snyk badge
		snykBadge := &common.Badge{
			Name: "snyk - known vulnerabilities",
		}
		snykBadge.Markdown = fmt.Sprintf("[![%s](https://snyk.io/test/npm/%s/badge.svg)](https://snyk.io/test/npm/%s)", snykBadge.Name, pj.Name, pj.Name)
		badges = append(badges, *snykBadge)
	}

	return badges, nil
}
