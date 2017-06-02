package badges

import (
	"fmt"
	"testing"
)

func TestNpmPackageContract(t *testing.T) {
	npmContract := &NpmBadgeContract{
		Path: "temp path",
	}
	badges, _ := npmContract.Badges()
	for i := 0; i < len(badges); i++ {
		fmt.Println(badges[i].Markdown)
	}
}
