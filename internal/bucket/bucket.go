package bucket

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/utils"
	"strings"
)

type Bucket struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
	Path   string `json:"path"`
}

func (b Bucket) Id() string {
	return fmt.Sprintf("%s:%s", b.Owner, b.Repo)
}

func (b Bucket) String() string {
	return fmt.Sprintf("%s@%s %s", b.Id(), b.Branch, b.Path)
}

func NewBucket(bucketExpression string, path string) *Bucket {
	var owner string
	var repo string
	var branch string

	if strings.Index(bucketExpression, "@") > 0 {
		var prefix string
		_ = utils.Unpack(strings.Split(bucketExpression, "@"), &prefix, &branch)
		_ = utils.Unpack(strings.Split(prefix, ":"), &owner, &repo)
	} else {
		branch = "main"
		_ = utils.Unpack(strings.Split(bucketExpression, ":"), &owner, &repo)
	}
	return &Bucket{
		Owner:  owner,
		Repo:   repo,
		Branch: branch,
		Path:   path,
	}
}
