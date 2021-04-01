package list

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/internal/section"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"path/filepath"
	"sort"
	"strings"
)

type Options struct {
	SectionName string
}

func NewCmdList(factory *cmdutil.Factory) *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			//client := factory.Client
			executableDir := filepath.Dir(factory.Executable)
			bucketRepository := factory.BucketRepository

			//err := cmdutil.IsExceededRateLimit(client)
			//if err != nil {
			//	return err
			//}

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

			output := make([]string, 0)
			if opts.SectionName != "" {
				if _, contain := sections[opts.SectionName]; !contain {
					return fmt.Errorf("invalid sectionObj name (%s)", opts.SectionName)
				}

				sectionObj := sections[opts.SectionName]

				output = append(output, fmt.Sprintf(sectionObj.Name))

				groupNames := sortedGroupNames(sectionObj.Groups)
				for idx, groupName := range groupNames {
					if idx < len(groupNames)-1 {
						output = append(output, fmt.Sprintf("├── %s", groupName))
					} else {
						output = append(output, fmt.Sprintf("└── %s", groupName))
					}
				}

			} else {
				sectionNames := sortedSectionNames(sections)
				for sectionIdx, sectionName := range sectionNames {
					if sectionIdx < len(sections)-1 {
						output = append(output, fmt.Sprintf("├── %s", sectionName))
					} else {
						output = append(output, fmt.Sprintf("└── %s", sectionName))
					}

					sectionObj := sections[sectionName]
					groupNames := sortedGroupNames(sectionObj.Groups)

					for groupIdx, groupName := range groupNames {
						var prefix string
						if sectionIdx < len(sections)-1 {
							prefix = "│  "
						} else {
							prefix = "    "
						}

						if groupIdx < len(sectionObj.Groups)-1 {
							output = append(output, fmt.Sprintf("%s├── %s", prefix, groupName))
						} else {
							output = append(output, fmt.Sprintf("%s└── %s", prefix, groupName))
						}
						groupIdx += 1
					}
				}
			}
			fmt.Print(strings.Join(output, "\n"))
			return nil
		},
	}

	cmd.Flags().StringVarP(
		&opts.SectionName, "section", "s", "", "Section Name",
	)
	return cmd
}

func sortedSectionNames(sections map[string]*section.Section) []string {
	idx := 0
	sectionNames := make([]string, len(sections))

	for sectionName := range sections {
		sectionNames[idx] = sectionName
		idx += 1
	}
	sort.Strings(sectionNames)
	return sectionNames
}

func sortedGroupNames(groups map[string]section.Group) []string {
	idx := 0
	groupNames := make([]string, len(groups))

	for groupName := range groups {
		groupNames[idx] = groupName
		idx += 1
	}
	sort.Strings(groupNames)
	return groupNames
}
