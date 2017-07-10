package downloader

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGitDownloader(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting current directory: ", err)
	}
	dir, err := ioutil.TempDir(wd, "data")
	if err != nil {
		log.Fatalln("Error creating temporary folder: ", err)
	}
	defer os.RemoveAll(dir)

	downloader := &GitDownloader{
		DownloaderOptions{
			Path:   dir,
			Remote: "https://github.com/scriptnull/ds.git",
			Type:   "git",
		},
	}

	err = downloader.Download()
	if err != nil {
		log.Fatalln("Expected error to be not null", err)
	}
}
