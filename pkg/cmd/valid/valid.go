package valid

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

var IssueCount = 0

func NewCmdValid() *cobra.Command {
	cmd := &cobra.Command{
		Use: "valid",
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

			err = validFile(cmd, bytes)
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

func validFile(cmd *cobra.Command, content []byte) error {
	m := make(map[string]interface{})
	err := yaml.Unmarshal(content, &m)

	if err != nil {
		return err
	}

	checkSectionInfo(cmd, m)

	if !isContain(m, "content") {
		IssueCount++
		cmd.Println("invalid: require 'content' field")
	} else {
		checkGroups(cmd, m["content"].([]interface{}))
	}
	return nil
}

func checkSectionInfo(cmd *cobra.Command, m map[string]interface{}) {
	if !isContain(m, "name") {
		IssueCount++
		cmd.Println("invalid: require 'name' field")
	}

	if !isContain(m, "source") {
		IssueCount++
		cmd.Println("invalid: require 'source' field")
	}
}

func checkGroups(cmd *cobra.Command, contents []interface{}) {
	for idx, v := range contents {
		checkGroup(cmd, idx, v)
	}
}

func checkGroup(cmd *cobra.Command, idx int, v interface{}) {
	m := v.(map[string]interface{})

	if !isContain(m, "name") {
		IssueCount++
		cmd.Printf("invalid group (index: %d): require 'name' field\n", idx)
	}

	if !isContain(m, "document") {
		IssueCount++
		cmd.Printf("invalid group (index: %d): require 'document' field\n", idx)
	}

	containJava := isContain(m, "java")
	containKotlin := isContain(m, "kotlin")

	if !containJava && !containKotlin {
		IssueCount++
		cmd.Printf("invalid group (index: %d): require 'java' or 'kotlin' field\n", idx)
	} else {
		if isContain(m, "java") {
			javaDeps := m["java"].([]interface{})
			for depIdx, dep := range javaDeps {
				checkDependency(cmd, idx, depIdx, dep)
			}
		}
		if isContain(m, "kotlin") {
			kotlinDeps := m["kotlin"].([]interface{})
			for depIdx, dep := range kotlinDeps {
				checkDependency(cmd, idx, depIdx, dep)
			}
		}
	}
}

func checkDependency(cmd *cobra.Command, groupIdx int, idx int, dep interface{}) {
	switch dep.(type) {
	case string:
		// Ignore
	case map[string]interface{}:
		m := dep.(map[string]interface{})

		if !isContain(m, "type") {
			IssueCount++
			cmd.Printf("invalid dependency (index: %d) in group (index: %d): require 'type' field\n", groupIdx, idx)
		}

		if !isContain(m, "content") {
			IssueCount++
			cmd.Printf("invalid dependency (index: %d) in group (index: %d): require 'content' field\n", groupIdx, idx)
		}
	default:
		IssueCount++
		cmd.Printf("invalid dependency (index: %s) in group (index: %d): Invalid format\n", groupIdx, idx)
	}
}

func isContain(m map[string]interface{}, key string) bool {
	_, contain := m[key]
	return contain
}
