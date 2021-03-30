package api

import (
	"bytes"
	"fmt"
	"strings"
)

type Trees struct {
	Sha       string `json:"sha"`
	Url       string `json:"url"`
	Tree      []Node `json:"tree"`
	Truncated bool   `json:"truncated"`
}

type Node struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	Sha  string `json:"sha"`
	Url  string `json:"url"`
}

func (tree Trees) FilterPath(path string) []Node {
	newNodes := make([]Node, 0)
	for _, node := range tree.Tree {
		if strings.HasPrefix(node.Path, path) {
			newNodes = append(newNodes, node)
		}
	}
	return newNodes
}

func GetTrees(client *Client, owner string, repo string, treeSha string, recursive bool) (*Trees, error) {
	if treeSha == "" {
		treeSha = "HEAD"
	}

	path := fmt.Sprintf("repos/%s/%s/git/trees/%s", owner, repo, treeSha)
	if recursive {
		path += "?recursive=1"
	}

	var trees Trees

	r := bytes.NewReader([]byte(`{}`))
	err := client.REST("https://api.github.com/", "GET", path, r, &trees)
	if err != nil {
		return nil, err
	}
	return &trees, nil
}
