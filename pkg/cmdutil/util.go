package cmdutil

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
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
