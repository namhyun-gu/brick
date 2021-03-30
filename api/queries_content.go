package api

import (
	"fmt"
)

func GetRawContent(client *Client, owner string, repo string, branch string, filepath string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", owner, repo, branch, filepath)
	return client.GET("https://raw.githubusercontent.com/", path)
}
