package api

import (
	"encoding/json"
	"fmt"
)

type Branch struct {
	Name string
	Sha  string
	Url  string
}

func GetBranch(client *Client, owner, repo, branch string) (*Branch, error) {
	path := fmt.Sprintf("repos/%s/%s/branches/%s", owner, repo, branch)

	content, err := client.GET("https://api.github.com/", path)
	if err != nil {
		return nil, err
	}

	return parseBranch(content)
}

func parseBranch(content []byte) (*Branch, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}

	name := m["name"].(string)
	commit := m["commit"].(map[string]interface{})
	commitInfo := commit["commit"].(map[string]interface{})
	tree := commitInfo["tree"].(map[string]interface{})

	sha := tree["sha"].(string)
	url := tree["url"].(string)

	return &Branch{
		Name: name,
		Sha:  sha,
		Url:  url,
	}, nil
}
