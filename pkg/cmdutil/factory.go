package cmdutil

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/pkg/iostreams"
)

type Factory struct {
	IO *iostreams.IOStreams

	Client           *api.Client
	BucketRepository *bucket.Repository

	Executable string
}
