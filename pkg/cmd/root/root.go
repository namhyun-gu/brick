package root

import (
	docCmd "github.com/namhyun-gu/brick/pkg/cmd/doc"
	getCmd "github.com/namhyun-gu/brick/pkg/cmd/get"
	listCmd "github.com/namhyun-gu/brick/pkg/cmd/list"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "brick",
		Short:   "Compose latest android dependencies",
		Version: "0.1.0",
	}

	cmd.AddCommand(getCmd.NewCmdGet(factory))
	cmd.AddCommand(docCmd.NewCmdDoc(factory))
	cmd.AddCommand(listCmd.NewCmdList(factory))
	return cmd
}
