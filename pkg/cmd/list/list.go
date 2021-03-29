package list

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/resource"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type ListOptions struct {
	SectionName string
}

func NewCmdList(factory *cmdutil.Factory) *cobra.Command {
	opts := &ListOptions{}

	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			output := make([]string, 0)
			if opts.SectionName != "" {
				if _, contain := sections[opts.SectionName]; !contain {
					return fmt.Errorf("invalid section name (%s)", opts.SectionName)
				}

				section := sections[opts.SectionName]

				output = append(output, fmt.Sprintf(section.Name))

				groupNames := sortedGroupNames(section.Groups)
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

					section := sections[sectionName]
					groupNames := sortedGroupNames(section.Groups)

					for groupIdx, groupName := range groupNames {
						var prefix string
						if sectionIdx < len(sections)-1 {
							prefix = "│  "
						} else {
							prefix = "    "
						}

						if groupIdx < len(section.Groups)-1 {
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

func sortedSectionNames(sections map[string]resource.Section) []string {
	idx := 0
	sectionNames := make([]string, len(sections))

	for sectionName := range sections {
		sectionNames[idx] = sectionName
		idx += 1
	}
	sort.Strings(sectionNames)
	return sectionNames
}

func sortedGroupNames(groups map[string]resource.Group) []string {
	idx := 0
	groupNames := make([]string, len(groups))

	for groupName := range groups {
		groupNames[idx] = groupName
		idx += 1
	}
	sort.Strings(groupNames)
	return groupNames
}
