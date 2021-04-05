package factory

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/internal/cache"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/namhyun-gu/brick/pkg/iostreams"
	"os"
	"path/filepath"
)

func New() (*cmdutil.Factory, error) {
	io := iostreams.System()

	appExecutable := "brick"
	if exe, err := os.Executable(); err == nil {
		appExecutable = exe
	}

	client := api.NewClient()

	var bucketCache cache.Cache
	bucketCache = cache.NewFileCache(filepath.Join(filepath.Dir(appExecutable), "buckets.json"))
	bucketRepository := bucket.NewBucketRepository(&bucketCache)

	return &cmdutil.Factory{
		IO:               io,
		Client:           client,
		BucketRepository: bucketRepository,
		Executable:       appExecutable,
	}, nil
}
