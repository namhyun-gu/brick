package cmdutil

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/bucket"
)

func IsExceededRateLimit(client *api.Client) error {
	limit, err := api.GetRateLimit(client)
	if err != nil {
		return err
	}

	if limit.Rate.Remaining == 0 {
		return fmt.Errorf("github API limit exceeded (limit: %d, reset: %d)", limit.Rate.Limit, limit.Rate.Reset)
	}
	return nil
}

func UpdateBucketCache(client *api.Client, repository *bucket.Repository, executableDir string) error {
	buckets, err := repository.Read()
	if err != nil {
		return err
	}

	for _, bucketObj := range buckets {
		fmt.Printf("Update %s...", bucketObj.Id())
		err := bucket.WriteCache(client, executableDir, bucketObj)
		if err != nil {
			return err
		}
		fmt.Println("Done")
	}
	return nil
}
