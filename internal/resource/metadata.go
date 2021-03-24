package resource

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/namhyun-gu/brick/api"
	"io/ioutil"
	"net/http"
	"strings"
)

type LibraryMetadata struct {
	xml.Name   `xml:"metadata"`
	GroupId    string             `xml:"groupId"`
	ArtifactId string             `xml:"artifactId"`
	Versions   LibraryVersionInfo `xml:"versioning"`
}

type LibraryVersionInfo struct {
	Latest      string         `xml:"latest"`
	Release     string         `xml:"release"`
	Versions    LibraryVersion `xml:"versions"`
	LastUpdated string         `xml:"lastUpdated"`
}

type LibraryVersion struct {
	Version []string `xml:"version"`
}

func GetSources(owner string, repo string) (map[string]string, error) {
	content, err := api.GetContents(owner, repo, "data/sources.json")
	if err != nil {
		return nil, err
	}

	rawContent, err := api.GetRawContent(content.DownloadUrl)
	if err != nil {
		return nil, err
	}

	if rawContent == nil {
		return nil, fmt.Errorf("raw content is nil")
	}
	return parseSources(rawContent)
}

func parseSources(content []byte) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func FetchMetadata(groupId string, artifactId string, source string) (*LibraryMetadata, error) {
	url := getMetadataUrl(groupId, artifactId, source)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed fetch metadata (status code: %d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	metadata, err := parseMetadata(body)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func parseMetadata(body []byte) (*LibraryMetadata, error) {
	var metadata LibraryMetadata
	err := xml.Unmarshal(body, &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func getMetadataUrl(groupId string, artifactId string, source string) string {
	return fmt.Sprintf(
		"%s/%s/%s/maven-metadata.xml",
		source,
		strings.ReplaceAll(groupId, ".", "/"),
		artifactId,
	)
}
