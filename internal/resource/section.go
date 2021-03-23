package resource

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Section struct {
	Name   string
	Source string
	Groups map[string]Group
}

type Group struct {
	Name     string
	Document string
	Java     []Dependency
	Kotlin   []Dependency
}

type Dependency struct {
	Configuration string
	Content       string
}

func GetSections(rootPath string) (map[string]Section, error) {
	contentsPath := filepath.Join(rootPath, "./data/sections")
	files, err := ioutil.ReadDir(contentsPath)
	if err != nil {
		return nil, err
	}

	var sections = make(map[string]Section)
	for _, file := range files {
		path := filepath.Join(contentsPath, file.Name())
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		section := parseSection(content)
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
	sectionSource := fmt.Sprintf("%v", m["source"])
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

	javaSlice := groupItem["java"].([]interface{})
	java := mapToDependency(javaSlice)

	var kotlin = make([]Dependency, 0)
	if _, ok := groupItem["kotlin"]; ok {
		kotlinSlice := groupItem["kotlin"].([]interface{})
		kotlin = mapToDependency(kotlinSlice)
	}

	return Group{
		Name:     name,
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
		return Dependency{
			Configuration: m["type"].(string),
			Content:       m["content"].(string),
		}
	}
}
