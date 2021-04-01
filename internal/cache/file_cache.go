package cache

import (
	"io/ioutil"
	"os"
)

type FileCache struct {
	Path string
}

func NewFileCache(path string) *FileCache {
	return &FileCache{
		Path: path,
	}
}

func (c FileCache) Write(content []byte) error {
	return ioutil.WriteFile(c.Path, content, 0666)
}

func (c FileCache) Read() ([]byte, error) {
	return ioutil.ReadFile(c.Path)
}

func (c FileCache) Exist() bool {
	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		return false
	}
	return true
}
