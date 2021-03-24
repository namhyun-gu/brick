package resource

import (
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"gopkg.in/yaml.v3"
	"log"
)

type Section struct {
	Name   string
	Source string
	Groups map[string]Group
}

type Group struct {
	Name     string
	Source   string
	Document string
	Java     []Dependency
	Kotlin   []Dependency
}

type Dependency struct {
	Configuration string
	Content       string
	Ignore        bool
}

func GetSections(owner string, repo string) (map[string]Section, error) {
	trees, err := api.GetTrees(owner, repo, "", true)
	if err != nil {
		return nil, err
	}

	var sections = make(map[string]Section)

	sectionNodes := trees.FilterPath("data/sections/")
	for _, sectionNode := range sectionNodes {
		content, err := api.GetContents(owner, repo, sectionNode.Path)
		if err != nil {
			return nil, err
		}

		rawContent, err := api.GetRawContent(content.DownloadUrl)
		if err != nil {
			return nil, err
		}

		if rawContent == nil {
			return nil, fmt.Errorf("raw content is nil")
		}

		section := parseSection(rawContent)
		sections[section.Name] = section
	}
	return sections, nil
}

func parseSection(content []byte) Section {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(content, &m)

	if err != nil {
		log.Fatal(err)
	}

	sectionName := fmt.Sprintf("%v", m["name"])

	var sectionSource = ""
	if _, contain := m["source"]; contain {
		sectionSource = fmt.Sprintf("%v", m["source"])
	}
	groupSlice := m["content"].([]interface{})
	groups := toGroupMap(groupSlice)

	return Section{
		Name:   sectionName,
		Source: sectionSource,
		Groups: groups,
	}
}

func toGroupMap(slice []interface{}) map[string]Group {
	m := make(map[string]Group)
	for _, v := range slice {
		group := parseGroup(v)
		m[group.Name] = group
	}
	return m
}

func parseGroup(group interface{}) Group {
	groupItem := group.(map[string]interface{})
	name := groupItem["name"].(string)
	document := groupItem["document"].(string)

	var source = ""
	if _, contain := groupItem["source"]; contain {
		source = groupItem["source"].(string)
	}

	javaSlice := groupItem["java"].([]interface{})
	java := mapToDependency(javaSlice)

	var kotlin = make([]Dependency, 0)
	if _, ok := groupItem["kotlin"]; ok {
		kotlinSlice := groupItem["kotlin"].([]interface{})
		kotlin = mapToDependency(kotlinSlice)
	}

	return Group{
		Name:     name,
		Source:   source,
		Document: document,
		Java:     java,
		Kotlin:   kotlin,
	}
}

func mapToDependency(slice []interface{}) []Dependency {
	temp := make([]Dependency, len(slice))
	for i, v := range slice {
		temp[i] = parseDependency(v)
	}
	return temp
}

func parseDependency(dep interface{}) Dependency {
	switch dep.(type) {
	case string:
		return Dependency{
			Configuration: "implementation",
			Content:       dep.(string),
		}
	default:
		m := dep.(map[string]interface{})
		ignore := false
		if _, contain := m["ignore"]; contain {
			if m["ignore"].(bool) {
				ignore = true
			}
		}

		configuration := "implementation"
		if _, contain := m["type"]; contain {
			configuration = m["type"].(string)
		}

		return Dependency{
			Configuration: configuration,
			Content:       m["content"].(string),
			Ignore:        ignore,
		}
	}
}
