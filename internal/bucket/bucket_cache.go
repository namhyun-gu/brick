package bucket

import (
	"github.com/namhyun-gu/brick/api"
	"github.com/namhyun-gu/brick/internal/cache"
	"github.com/namhyun-gu/brick/internal/section"
	"github.com/namhyun-gu/brick/internal/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func WriteCache(client *api.Client, path string, bucket *Bucket) error {
	trees, err := api.GetTrees(client, bucket.Owner, bucket.Repo, "", true)
	if err != nil {
		return err
	}

	sectionPaths := make([]string, 0)
	sectionNodes := trees.FilterPath(bucket.Path)
	for _, sectionNode := range sectionNodes {
		sectionPaths = append(sectionPaths, sectionNode.Path)
	}

	cacheDir := cacheDirPath(path, bucket)
	if !utils.ExistFile(cacheDir) {
		err := os.MkdirAll(cacheDir, 0666)
		if err != nil {
			return err
		}
	}

	for _, path := range sectionPaths {
		raw, err := api.GetRawContent(client, bucket.Owner, bucket.Repo, bucket.Branch, path)
		if err != nil {
			return err
		}

		sectionFileName := strings.TrimPrefix(path, bucket.Path)
		sectionFilePath := filepath.Join(cacheDir, sectionFileName)

		newCache := cache.NewFileCache(sectionFilePath)
		err = newCache.Write(raw)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadCache(path string, bucket *Bucket) (map[string]*section.Section, error) {
	cacheDir := cacheDirPath(path, bucket)

	dirs, err := ioutil.ReadDir(cacheDir)
	if err != nil {
		return nil, err
	}

	sectionMap := make(map[string]*section.Section)
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}

		filename := filepath.Join(cacheDir, dir.Name())
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		s, err := section.ParseSection(file)
		if err != nil {
			return nil, err
		}
		sectionMap[s.Name] = s
	}
	return sectionMap, nil
}

func cacheDirPath(path string, bucket *Bucket) string {
	dirname := strings.Replace(bucket.Id(), ":", "/", 1)
	return filepath.Join(path, "cache", dirname)
}
