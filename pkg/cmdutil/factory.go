package cmdutil

import (
	"github.com/namhyun-gu/brick/api"
)

type Factory struct {
	Client *api.Client

	Executable string
}
