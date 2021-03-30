package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
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

func GetMetadata(client *Client, source string, groupId string, artifactId string) (*LibraryMetadata, error) {
	path := fmt.Sprintf("/%s/%s/maven-metadata.xml", strings.ReplaceAll(groupId, ".", "/"), artifactId)

	r := bytes.NewReader([]byte(`{}`))
	body, err := client.GET(source, path, r)

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
