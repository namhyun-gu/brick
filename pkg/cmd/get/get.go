package get

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/resource"
	"github.com/namhyun-gu/brick/internal/util"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"path/filepath"
	"strings"
)

type GetOptions struct {
	ProjectLanguage string
	GradleLanguage  string
}

type Argument struct {
	SectionName string
	GroupName   string
}

func (argument *Argument) IsValid(sections map[string]resource.Section, ) bool {
	if _, contain := sections[argument.SectionName]; !contain {
		return false
	}

	section := sections[argument.SectionName]

	if _, contain := section.Groups[argument.GroupName]; !contain {
		return false
	}
	return true
}

type FetchJob struct {
	SectionName   string
	GroupName     string
	GroupId       string
	ArtifactId    string
	Configuration string
	Source        string
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
			requests := make([]Argument, 0)

			for _, arg := range args {
				var sectionName, groupName string
				err := util.Unpack(strings.Split(arg, ":"), &sectionName, &groupName)
				if err != nil {
					return fmt.Errorf("invalid argument (arg: [%s])", arg)
				}

				requests = append(requests, Argument{
					SectionName: sectionName,
					GroupName:   groupName,
				})
			}

			executableDir := filepath.Dir(factory.Executable)
			sources, err := resource.GetSources(executableDir)
			if err != nil {
				return err
			}

			sections, err := resource.GetSections(executableDir)
			if err != nil {
				return err
			}

			return RunGet(opts, sources, sections, requests)
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

func RunGet(
	opts *GetOptions,
	sources map[string]string,
	sections map[string]resource.Section,
	arguments []Argument,
) error {
	validArguments := funk.Filter(arguments, func(argument Argument) bool {
		return argument.IsValid(sections)
	}).([]Argument)
	argumentMap := groupArguments(validArguments)
	jobs := make([]FetchJob, 0)

	for sectionName, groupNames := range argumentMap {
		section := sections[sectionName]
		source := getSource(sources, section.Source)

		for _, groupName := range groupNames {
			group := section.Groups[groupName]

			if opts.ProjectLanguage == "kotlin" && len(group.Kotlin) > 0 {
				newJobs := funk.Map(group.Kotlin, func(dependency resource.Dependency) FetchJob {
					return makeFetchJob(sectionName, groupName, source, dependency)
				}).([]FetchJob)

				jobs = append(jobs, newJobs...)
			} else {
				newJobs := funk.Map(group.Java, func(dependency resource.Dependency) FetchJob {
					return makeFetchJob(sectionName, groupName, source, dependency)
				}).([]FetchJob)

				jobs = append(jobs, newJobs...)
			}
		}
	}

	output := make([]string, 0)
	for _, job := range jobs {
		metadata, err := resource.FetchMetadata(
			job.GroupId,
			job.ArtifactId,
			job.Source,
		)

		if err != nil {
			return err
		}

		dependencyString := util.MakeDependencyString(
			job.Configuration,
			job.GroupId,
			job.ArtifactId,
			metadata.Versions.Latest,
			opts.GradleLanguage,
		)
		output = append(output, dependencyString)
	}

	fmt.Print(strings.Join(output, "\n"))
	return nil
}

func makeFetchJob(
	sectionName string,
	groupName string,
	source string,
	dependency resource.Dependency,
) FetchJob {
	var groupId, artifactId string
	_ = util.Unpack(strings.Split(dependency.Content, ":"), &groupId, &artifactId)

	return FetchJob{
		SectionName:   sectionName,
		GroupName:     groupName,
		GroupId:       groupId,
		ArtifactId:    artifactId,
		Configuration: dependency.Configuration,
		Source:        source,
	}
}

func groupArguments(arguments []Argument) map[string][]string {
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
