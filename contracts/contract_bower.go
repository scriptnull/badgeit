package contracts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/scriptnull/badgeit/common"
)

// BowerBadgeContract checks for npm badges
type BowerBadgeContract struct {
	Path string
}

// NewBowerBadgeContract returns contract for checking various badges related to npm
func NewBowerBadgeContract(path string) *BowerBadgeContract {
	return &BowerBadgeContract{
		Path: path,
	}
}

// Badges yields array of badges that satisfy the contract
func (contract *BowerBadgeContract) Badges() ([]common.Badge, error) {

	// read bower.json
	raw, err := ioutil.ReadFile(filepath.Join(contract.Path, "bower.json"))
	if err != nil {
		return []common.Badge{}, err
	}

	// parse json from bower.json
	bowerJSON := struct {
		Name     string
		Homepage string
	}{}
	err = json.Unmarshal(raw, &bowerJSON)
	if err != nil {
		return []common.Badge{}, err
	}

	var badges []common.Badge

	if len(bowerJSON.Name) > 0 {
		shieldBaseURL := "https://img.shields.io/bower"

		// version Badge
		normalVersion := &common.Badge{
			Name:     "bower version",
			ImageURL: fmt.Sprintf("%s/v/%s.svg", shieldBaseURL, bowerJSON.Name),
			LinkURL:  fmt.Sprintf("%s", bowerJSON.Homepage),
			Tags:     []string{"version"},
		}
		badges = append(badges, *normalVersion)

		// license badge
		licenseBadge := &common.Badge{
			Name:     "license badge",
			ImageURL: fmt.Sprintf("%s/l/%s.svg", shieldBaseURL, bowerJSON.Name),
			LinkURL:  fmt.Sprintf("%s", bowerJSON.Homepage),
			Tags:     []string{"license"},
		}
		badges = append(badges, *licenseBadge)
	}

	commonTags := []string{"bower", "package manager"}
	for i := 0; i < len(badges); i++ {
		badges[i].Group = "bower"
		badges[i].Tags = append(badges[i].Tags, commonTags...)
	}

	return badges, nil
}
