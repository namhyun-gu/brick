package valid

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Lang string

const (
	Java   Lang = "java"
	Kotlin      = "kotlin"
)

type Locator struct {
	Node  *yaml.Node
	Group string
}

func (loc *Locator) string() string {
	if loc.Group == "" {
		if loc.Node != nil {
			return fmt.Sprintf("(line: %d, col: %d)", loc.Node.Line, loc.Node.Column)
		} else {
			return "(root)"
		}
	} else {
		if loc.Node != nil {
			return fmt.Sprintf("%s (line: %d, col: %d)", loc.Group, loc.Node.Line, loc.Node.Column)
		} else {
			return fmt.Sprintf("'%s'", loc.Group)
		}
	}
}

var IssueCount = 0

func NewCmdValid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "valid",
		Short: "Validate configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			targetFile := args[0]
			targetFilePath := filepath.Join(wd, targetFile)

			if _, err := os.Stat(targetFilePath); os.IsNotExist(err) {
				return fmt.Errorf("%s not found", targetFile)
			}
			bytes, err := ioutil.ReadFile(targetFilePath)
			if err != nil {
				return err
			}

			err = validFile(bytes)
			if err != nil {
				return err
			}

			if IssueCount == 0 {
				cmd.Println("No issue found")
				return nil
			} else {
				return fmt.Errorf("%d issues found", IssueCount)
			}
		},
	}
	return cmd
}

type Root struct {
	Name    yaml.Node   `yaml:"name"`
	Source  yaml.Node   `yaml:"source"`
	Content []yaml.Node `yaml:"content"`
}

type Group struct {
	Name     yaml.Node   `yaml:"name"`
	Source   string      `yaml:"source"`
	Document yaml.Node   `yaml:"document"`
	Java     []yaml.Node `yaml:"java"`
	Kotlin   []yaml.Node `yaml:"kotlin"`
}

type Dependency struct {
	Type    yaml.Node `yaml:"type"`
	Content yaml.Node `yaml:"content"`
}

func validFile(content []byte) error {
	var root Root
	err := yaml.Unmarshal(content, &root)

	if err != nil {
		return err
	}
	return validRoot(root)
}

func validRoot(root Root) error {
	if root.Source.Value == "" {
		invalid("require 'source' field", Locator{})
	}

	if root.Name.Value == "" {
		invalid("require 'name' field", Locator{})
	}

	if root.Content == nil {
		invalid("require 'content' field", Locator{})
	}

	err := validContents(root.Content, root.Source.Value)
	if err != nil {
		return err
	}
	return nil
}

func validContents(contents []yaml.Node, rootSource string) error {
	for _, content := range contents {
		var group Group
		err := content.Decode(&group)
		if err != nil {
			return err
		}
		err = validGroup(content, group, rootSource)
		if err != nil {
			return err
		}
	}
	return nil
}

func validGroup(node yaml.Node, group Group, rootSource string) error {
	groupName := group.Name.Value
	if group.Name.Value == "" {
		invalid("require 'name' field", Locator{Group: groupName, Node: &node})
	}
	if group.Document.Value == "" {
		invalid("require 'document' field", Locator{Group: groupName, Node: &node})
	}
	if group.Java == nil && group.Kotlin == nil {
		invalid("require 'java' or 'kotlin' field", Locator{Group: groupName, Node: &node})
	}

	var source string
	if group.Source != "" {
		source = group.Source
	} else {
		source = rootSource
	}

	if group.Java != nil {
		return validDependencies(Java, groupName, group.Java, source)
	}
	if group.Kotlin != nil {
		return validDependencies(Kotlin, groupName, group.Kotlin, source)
	}
	return nil
}

func validDependencies(lang Lang, group string, depNodes []yaml.Node, source string) error {
	group = fmt.Sprintf("%s (%s)", group, lang)
	for _, depNode := range depNodes {
		if depNode.Kind == yaml.ScalarNode {
			if depNode.Value == "" {
				invalid("require content value", Locator{Group: group, Node: &depNode})
			}
			if depNode.Value != "" {
				err := validLibrary(source, depNode.Value)
				if err != nil {
					if source == "" {
						source = "null"
					}
					invalid(fmt.Sprintf("unable '%s' (source: %s)", depNode.Value, source), Locator{
						Node:  &depNode,
						Group: group,
					})
				}
			}
		} else {
			var dep Dependency
			err := depNode.Decode(&dep)
			if err != nil {
				return err
			}
			validDependency(source, group, depNode, dep)
		}
	}
	return nil
}

func validDependency(source string, group string, node yaml.Node, dep Dependency) {
	if dep.Type.Value == "" {
		invalid("require 'type' field", Locator{Group: group, Node: &node})
	}
	if dep.Content.Value == "" {
		invalid("require 'content' field", Locator{Group: group, Node: &node})
	}
	if dep.Content.Value != "" {
		err := validLibrary(source, dep.Content.Value)
		if err != nil {
			if source == "" {
				source = "null"
			}
			invalid(fmt.Sprintf("unable '%s' (source: %s)", dep.Content.Value, source), Locator{Group: group, Node: &node})
		}
	}
}

func invalid(message string, locator Locator) {
	IssueCount++
	fmt.Printf("invalid: %s in %s\n", message, locator.string())
}

func validLibrary(source string, library string) error {
	source = api.MavenSource()(source)

	var groupId, artifactId string
	err := utils.Unpack(strings.Split(library, ":"), &groupId, &artifactId)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/%s/%s/maven-metadata.xml", strings.ReplaceAll(groupId, ".", "/"), artifactId)
	client := api.NewClient()

	_, err = client.GET(source, path)
	if err != nil {
		return err
	}
	return nil
}
