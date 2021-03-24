package utils

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/browser"
)

func Unpack(slice []string, vars ...*string) error {
	if len(slice) < len(vars) {
		return fmt.Errorf("invalid slice size (slice: %d, vargs: %d)", len(slice), len(vars))
	}
	for i, value := range slice {
		*vars[i] = value
	}
	return nil
}

func MakeDependencyString(
	configuration string,
	groupId string,
	artifactId string,
	version string,
	gradleLanguage string,
) string {
	var dependencyNotation string
	if version != "" {
		dependencyNotation = fmt.Sprintf("%s:%s:%s", groupId, artifactId, version)
	} else {
		dependencyNotation = fmt.Sprintf("%s:%s", groupId, artifactId)
	}

	if configuration == "platform" {
		if gradleLanguage == "groovy" {
			return fmt.Sprintf("%s platform(\"%s\")", "implementation", dependencyNotation)
		} else if gradleLanguage == "kotlin" {
			return fmt.Sprintf("%s(platform(\"%s\"))", "implementation", dependencyNotation)
		}
	} else {
		if gradleLanguage == "groovy" {
			return fmt.Sprintf("%s \"%s\"", configuration, dependencyNotation)
		} else if gradleLanguage == "kotlin" {
			return fmt.Sprintf("%s(\"%s\")", configuration, dependencyNotation)
		}
	}
	return ""
}

func OpenInBrowser(url string) error {
	browserCmd, err := browser.Command(url)
	if err != nil {
		return err
	}
	return browserCmd.Run()
}
