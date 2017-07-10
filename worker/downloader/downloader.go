package downloader

type Downloader interface {
	Download() error
}

type DownloaderOptions struct {
	Type   string
	Remote string
	Path   string
}

func NewDownloader(opts DownloaderOptions) Downloader {
	switch opts.Type {
	case "git":
		return &GitDownloader{opts}
	case "curl":
		return &CurlDownloader{opts}
	}
	return nil
}
