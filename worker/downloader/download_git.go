package downloader

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

type GitDownloader struct {
	DownloaderOptions
}

func (down *GitDownloader) Download() error {
	_, err := git.PlainClone(down.Path, false, &git.CloneOptions{
		URL:      down.Remote,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}
