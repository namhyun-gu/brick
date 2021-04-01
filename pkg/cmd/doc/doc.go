package doc

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/internal/section"
	"github.com/namhyun-gu/brick/internal/utils"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"path/filepath"
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

			//client := factory.Client
			executableDir := filepath.Dir(factory.Executable)
			bucketRepository := factory.BucketRepository

			buckets, err := bucketRepository.Read()
			if err != nil {
				return err
			}

			sections := make(map[string]*section.Section)
			for _, bucketObj := range buckets {
				bucketSections, err := bucket.ReadCache(executableDir, bucketObj)
				if err != nil {
					return err
				}

				sections = section.ConcatSectionMap(sections, bucketSections)
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
