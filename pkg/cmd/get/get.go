package get

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
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
	Dependency  api.Dependency
	Source      string
}

type FetchJobResult struct {
	SectionName   string
	GroupName     string
	Configuration string
	Metadata      *api.LibraryMetadata
}

func NewCmdGet(factory *cmdutil.Factory) *cobra.Command {
	opts := &GetOptions{}

	cmd := &cobra.Command{
		Use:  "get [{section}:{group}]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			arguments := make([]*utils.Argument, 0)
			client := factory.Client

			for _, arg := range args {
				argument, err := utils.ParseArgument(arg)
				if err != nil {
					return err
				}

				arguments = append(arguments, argument)
			}
			return getRun(client, arguments, opts)
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
	client *api.Client,
	arguments []*utils.Argument,
	opts *GetOptions,
) error {
	err := cmdutil.IsExceededRateLimit(client)
	if err != nil {
		return err
	}

	sections, err := api.GetSections(client, "namhyun-gu", "brick", "main")
	if err != nil {
		return err
	}

	argumentMap := groupArguments(filterValidArguments(arguments, sections))
	fetchJobs := makeFetchJobs(opts, sections, argumentMap)
	output, err := runFetchJobs(client, opts, fetchJobs)
	if err != nil {
		return err
	}
	fmt.Print(strings.Join(output, "\n"))
	return nil
}

func makeFetchJobs(
	opts *GetOptions,
	sections map[string]*api.Section,
	argumentMap map[string][]string,
) []FetchJob {
	jobs := make([]FetchJob, 0)
	mavenSource := api.MavenSource()

	for sectionName, groupNames := range argumentMap {
		section := sections[sectionName]
		source := mavenSource(section.Source)

		for _, groupName := range groupNames {
			group := section.Groups[groupName]

			var groupSource string
			if group.Source != "" {
				groupSource = mavenSource(group.Source)
			} else {
				groupSource = source
			}

			var dependencies []api.Dependency

			if opts.ProjectLanguage == "kotlin" && len(group.Kotlin) > 0 {
				dependencies = group.Kotlin
			} else {
				dependencies = group.Java
			}

			newJobs := funk.Map(dependencies, func(dependency api.Dependency) FetchJob {
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
	return jobs
}

func runFetchJobs(
	client *api.Client,
	opts *GetOptions,
	fetchJobs []FetchJob,
) ([]string, error) {
	output := make([]string, 0)
	for _, job := range fetchJobs {
		var groupId, artifactId string
		err := utils.Unpack(strings.Split(job.Dependency.Content, ":"), &groupId, &artifactId)
		if err != nil {
			return nil, err
		}

		latestVersion := ""
		if !job.Dependency.Ignore {
			metadata, err := api.GetMetadata(
				client,
				job.Source,
				groupId,
				artifactId,
			)

			if err != nil {
				return nil, err
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
	return output, nil
}

func filterValidArguments(arguments []*utils.Argument, sections map[string]*api.Section) []*utils.Argument {
	return funk.Filter(arguments, func(argument *utils.Argument) bool {
		return argument.IsValid(sections)
	}).([]*utils.Argument)
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
