package section

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

func ParseSection(content []byte) (*Section, error) {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(content, &m)

	if err != nil {
		return nil, err
	}

	sectionName := fmt.Sprintf("%v", m["name"])

	var sectionSource = ""
	if _, contain := m["source"]; contain {
		sectionSource = fmt.Sprintf("%v", m["source"])
	}
	groupSlice := m["content"].([]interface{})
	groups := toGroupMap(groupSlice)

	return &Section{
		Name:   sectionName,
		Source: sectionSource,
		Groups: groups,
	}, nil
}

func ConcatSectionMap(a, b map[string]*Section) map[string]*Section {
	newMap := make(map[string]*Section)
	for key, s := range a {
		newMap[key] = s
	}
	for key, s := range b {
		newMap[key] = s
	}
	return newMap
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
