package contracts

import "github.com/scriptnull/badgeit/common"

// PossibleBadges returns the possible badges detected for a path
func PossibleBadges(path string) []common.Badge {
	// Check Contract aggreement and obtain eligible badges
	var badges []common.Badge

	npmBadges, err := NewNpmBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, npmBadges...)
	}
	githubBadges, err := NewGithubBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, githubBadges...)
	}
	gitterBadges, err := NewGitterBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, gitterBadges...)
	}
	bowerBadges, err := NewBowerBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, bowerBadges...)
	}
	travisBadges, err := NewTravisBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, travisBadges...)
	}
	circleBadges, err := NewCircleBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, circleBadges...)
	}
	semaphoreBadges, err := NewSemaphoreBadgeContract(path).Badges()
	if err == nil {
		badges = append(badges, semaphoreBadges...)
	}

	return badges
}
