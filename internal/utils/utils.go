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
	dependencyNotation := fmt.Sprintf("%s:%s:%s", groupId, artifactId, version)
	if gradleLanguage == "groovy" {
		return fmt.Sprintf("%s \"%s\"", configuration, dependencyNotation)
	} else if gradleLanguage == "kotlin" {
		return fmt.Sprintf("%s(\"%s\")", configuration, dependencyNotation)
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
