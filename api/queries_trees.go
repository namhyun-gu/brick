package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetTrees(owner string, repo string, treeSha string, recursive bool) (*Trees, error) {
	if treeSha == "" {
		treeSha = "HEAD"
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s", owner, repo, treeSha)
	if recursive {
		url += "?recursive=1"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed get trees (code: %d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var trees Trees
	err = json.Unmarshal(body, &trees)
	if err != nil {
		return nil, err
	}
	return &trees, nil
}
