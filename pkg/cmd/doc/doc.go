package doc

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/resource"
	"github.com/namhyun-gu/brick/internal/utils"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdDoc(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "doc [{section}:{group}]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argument, err := utils.ParseArgument(args[0])
			if err != nil {
				return err
			}

			limit, err := api.GetRateLimit()
			if err != nil {
				return err
			}

			if limit.Rate.Remaining == 0 {
				return fmt.Errorf("github API limit exceeded (limit: %d, reset: %d)", limit.Rate.Limit, limit.Rate.Reset)
			}

			sections, err := resource.GetSections("namhyun-gu", "brick")
			if err != nil {
				return err
			}

			if !argument.IsValid(sections) {
				return fmt.Errorf("invalid argument (arg: %s)", argument.ToString())
			}

			section := sections[argument.SectionName]
			group := section.Groups[argument.GroupName]

			fmt.Printf("Opening %s in your browser.", group.Document)
			return utils.OpenInBrowser(group.Document)
		},
	}
	return cmd
}
