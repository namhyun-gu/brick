package utils

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/resource"
	"strings"
)

type Argument struct {
	SectionName string
	GroupName   string
}

func (argument *Argument) IsValid(sections map[string]resource.Section) bool {
	if _, contain := sections[argument.SectionName]; !contain {
		return false
	}

	section := sections[argument.SectionName]

	if _, contain := section.Groups[argument.GroupName]; !contain {
		return false
	}
	return true
}

func (argument *Argument) ToString() string {
	return fmt.Sprintf("%s:%s", argument.SectionName, argument.GroupName)
}

func ParseArgument(arg string) (*Argument, error) {
	var sectionName, groupName string
	err := Unpack(strings.Split(arg, ":"), &sectionName, &groupName)
	if err != nil {
		return nil, fmt.Errorf("invalid argument (arg: %s)", arg)
	}
	return &Argument{
		SectionName: sectionName,
		GroupName:   groupName,
	}, nil
}
