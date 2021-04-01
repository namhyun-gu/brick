package update

import (
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"path/filepath"
)

func NewCmdUpdate(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := factory.Client
			repository := factory.BucketRepository
			executableDir := filepath.Dir(factory.Executable)

			return cmdutil.UpdateBucketCache(client, repository, executableDir)
		},
	}

	return cmd
}
