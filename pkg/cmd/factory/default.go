package factory

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"os"
)

func New() *cmdutil.Factory {
	appExecutable := "brick"
	if exe, err := os.Executable(); err == nil {
		appExecutable = exe
	}

	return &cmdutil.Factory{
		Client:     api.NewClient(),
		Executable: appExecutable,
	}
}
