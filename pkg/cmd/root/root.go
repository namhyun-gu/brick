package root

import (
	bucketCmd "github.com/namhyun-gu/brick/pkg/cmd/bucket"
	docCmd "github.com/namhyun-gu/brick/pkg/cmd/doc"
	getCmd "github.com/namhyun-gu/brick/pkg/cmd/get"
	listCmd "github.com/namhyun-gu/brick/pkg/cmd/list"
	updateCmd "github.com/namhyun-gu/brick/pkg/cmd/update"
	validCmd "github.com/namhyun-gu/brick/pkg/cmd/valid"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "brick <command> <subcommand>",
		Short:   "Compose latest android dependencies",
		Version: "0.1.0",
	}

	cmd.AddCommand(getCmd.NewCmdGet(factory))
	cmd.AddCommand(docCmd.NewCmdDoc(factory))
	cmd.AddCommand(listCmd.NewCmdList(factory))
	cmd.AddCommand(updateCmd.NewCmdUpdate(factory))
	cmd.AddCommand(bucketCmd.NewCmdBucket(factory))
	cmd.AddCommand(validCmd.NewCmdValid())
	return cmd
}
