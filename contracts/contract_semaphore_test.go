package contracts

import (
	"testing"
)

func TestGenerateSemaphoreBadge(t *testing.T) {
	var testMap = map[string]bool{
		"[![semaphore master badge](https://semaphoreci.com/api/v1/argonlaser/badgeit-front/branches/master/badge.svg)](https://semaphoreci.com/argonlaser/badgeit-front)": false,
	}
	contract := NewSemaphoreBadgeContract("../samples/semaphore/valid")
	badges, err := contract.Badges()
	if err != nil {
		t.Error("Expected err to be not nil, but got", err)
	}

	for _, badge := range badges {
		testMap[badge.Markdown()] = true
	}

	for key, val := range testMap {
		if val == false {
			t.Error("Failing Badge generation for", key)
		}
	}
}

func TestInvalidSemaphoreBadge(t *testing.T) {
	contract := NewSemaphoreBadgeContract("../samples/semaphore/invalid")
	badges, err := contract.Badges()
	if err != nil {
		t.Error("Expected err to be not nil, but got", err)
	}

	if len(badges) != 0 {
		t.Errorf("Expected to not generate badges for semaphore disabled repos")
	}
}
