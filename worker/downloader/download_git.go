package downloader

import (
	"os/exec"
)

// GitDownloader downloads a git repository
type GitDownloader struct {
	DownloaderOptions
}

// Download clones the git repository
func (down *GitDownloader) Download() error {
	cmd := exec.Command("git", "clone", down.Remote)
	cmd.Dir = down.Path
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
