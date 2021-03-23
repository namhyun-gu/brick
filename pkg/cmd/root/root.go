package root

import (
	getCmd "github.com/namhyun-gu/brick/pkg/cmd/get"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "brick",
		Short:   "Compose latest android library dependencies",
		Version: "0.1.0",
	}

	cmd.AddCommand(getCmd.NewCmdGet(factory))
	
	return cmd
}
