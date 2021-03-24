package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type File struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int64  `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	Sha         string `json:"sha"`
	Url         string `json:"url"`
	GitUrl      string `json:"git_url"`
	HtmlUrl     string `json:"html_url"`
	DownloadUrl string `json:"download_url"`
}

func GetContents(owner string, repo string, path string) (*File, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed get contents (code: %d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var file File
	err = json.Unmarshal(body, &file)
	if err != nil {
		return nil, err
	}
	return &file, err
}

func GetRawContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed get raw content (code: %d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
