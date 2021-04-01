package cmdutil

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/bucket"
)

type Factory struct {
	Client           *api.Client
	BucketRepository *bucket.Repository

	Executable string
}
