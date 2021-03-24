package get

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/resource"
	"github.com/namhyun-gu/brick/internal/utils"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"strings"
)

type GetOptions struct {
	ProjectLanguage string
	GradleLanguage  string
}

type FetchJob struct {
	SectionName string
	GroupName   string
	Dependency  resource.Dependency
	Source      string
}

type FetchJobResult struct {
	SectionName   string
	GroupName     string
	Configuration string
	Metadata      *resource.LibraryMetadata
}

func NewCmdGet(factory *cmdutil.Factory) *cobra.Command {
	opts := &GetOptions{}

	cmd := &cobra.Command{
		Use:  "get [{section}:{group}]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			requests := make([]*utils.Argument, 0)

			for _, arg := range args {
				argument, err := utils.ParseArgument(arg)
				if err != nil {
					return err
				}

				requests = append(requests, argument)
			}

			sources, err := resource.GetSources("namhyun-gu", "brick")
			if err != nil {
				return err
			}

			sections, err := resource.GetSections("namhyun-gu", "brick")
			if err != nil {
				return err
			}

			return getRun(opts, sources, sections, requests)
		},
	}

	cmd.Flags().StringVarP(
		&opts.ProjectLanguage, "lang", "l", "kotlin", "Project Language (kotlin or java)",
	)

	cmd.Flags().StringVarP(
		&opts.GradleLanguage, "gradle", "g", "groovy", "Gradle Language (groovy or kotlin)",
	)
	return cmd
}

func getRun(
	opts *GetOptions,
	sources map[string]string,
	sections map[string]resource.Section,
	arguments []*utils.Argument,
) error {
	validArguments := funk.Filter(arguments, func(argument *utils.Argument) bool {
		return argument.IsValid(sections)
	}).([]*utils.Argument)
	argumentMap := groupArguments(validArguments)
	jobs := make([]FetchJob, 0)

	for sectionName, groupNames := range argumentMap {
		section := sections[sectionName]
		source := getSource(sources, section.Source)

		for _, groupName := range groupNames {
			group := section.Groups[groupName]

			var groupSource string
			if group.Source != "" {
				groupSource = getSource(sources, group.Source)
			} else {
				groupSource = source
			}

			var dependencies []resource.Dependency

			if opts.ProjectLanguage == "kotlin" && len(group.Kotlin) > 0 {
				dependencies = group.Kotlin
			} else {
				dependencies = group.Java
			}

			newJobs := funk.Map(dependencies, func(dependency resource.Dependency) FetchJob {
				return FetchJob{
					SectionName: sectionName,
					GroupName:   groupName,
					Dependency:  dependency,
					Source:      groupSource,
				}
			}).([]FetchJob)

			jobs = append(jobs, newJobs...)
		}
	}

	output := make([]string, 0)
	for _, job := range jobs {
		var groupId, artifactId string
		err := utils.Unpack(strings.Split(job.Dependency.Content, ":"), &groupId, &artifactId)
		if err != nil {
			return err
		}

		latestVersion := ""
		if !job.Dependency.Ignore {
			metadata, err := resource.FetchMetadata(
				groupId,
				artifactId,
				job.Source,
			)

			if err != nil {
				return err
			}

			latestVersion = metadata.Versions.Latest
		}

		dependencyString := utils.MakeDependencyString(
			job.Dependency.Configuration,
			groupId,
			artifactId,
			latestVersion,
			opts.GradleLanguage,
		)
		output = append(output, dependencyString)
	}

	fmt.Print(strings.Join(output, "\n"))
	return nil
}

func groupArguments(arguments []*utils.Argument) map[string][]string {
	m := make(map[string][]string)
	for _, argument := range arguments {
		if _, contain := m[argument.SectionName]; !contain {
			m[argument.SectionName] = make([]string, 0)
		}
		m[argument.SectionName] = append(m[argument.SectionName], argument.GroupName)
	}
	return m
}

func getSource(sources map[string]string, source string) string {
	if _, contain := sources[source]; !contain {
		return source
	}
	return sources[source]
}
