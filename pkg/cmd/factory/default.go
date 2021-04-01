package factory

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/internal/cache"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"os"
	"path/filepath"
)

var defaultBuckets = []*bucket.Bucket{
	bucket.NewBucket("namhyun-gu:brick@main", "data/"),
}

func New() (*cmdutil.Factory, error) {
	appExecutable := "brick"
	if exe, err := os.Executable(); err == nil {
		appExecutable = exe
	}

	client := api.NewClient()

	var bucketCache cache.Cache
	bucketCachePath := filepath.Join(filepath.Dir(appExecutable), "buckets.json")
	bucketCache = cache.NewFileCache(bucketCachePath)
	bucketRepository := bucket.NewBucketRepository(&bucketCache)

	if !bucketCache.Exist() {
		err := bucketRepository.SaveAll(defaultBuckets)
		if err != nil {
			return nil, err
		}

		err = cmdutil.UpdateBucketCache(client, bucketRepository, filepath.Dir(appExecutable))
		if err != nil {
			return nil, err
		}
	}

	return &cmdutil.Factory{
		Client:           client,
		BucketRepository: bucketRepository,
		Executable:       appExecutable,
	}, nil
}
