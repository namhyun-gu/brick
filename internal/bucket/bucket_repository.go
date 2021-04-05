package bucket

import (
	"encoding/json"
	"github.com/namhyun-gu/brick/internal/cache"
)

type Repository struct {
	cache *cache.Cache
}

func NewBucketRepository(cache *cache.Cache) *Repository {
	return &Repository{
		cache: cache,
	}
}

func (repo Repository) Save(newBucket *Bucket) error {
	buckets, err := repo.Read()
	if err != nil {
		return err
	}

	bucketMap := groupBuckets(buckets)

	newBucketId := newBucket.Id()
	bucketMap[newBucketId] = newBucket

	newBuckets := flatBucketMap(bucketMap)
	return repo.SaveAll(newBuckets)
}

func (repo Repository) SaveAll(buckets []*Bucket) error {
	jsonContent, err := json.Marshal(buckets)
	if err != nil {
		return err
	}
	return (*repo.cache).Write(jsonContent)
}

func (repo Repository) Remove(bucketId string) error {
	buckets, err := repo.Read()
	if err != nil {
		return err
	}

	bucketMap := groupBuckets(buckets)
	delete(bucketMap, bucketId)

	newBuckets := flatBucketMap(bucketMap)
	return repo.SaveAll(newBuckets)
}

func (repo Repository) Read() ([]*Bucket, error) {
	c := *repo.cache
	if !c.Exist() {
		return nil, nil
	}

	bytes, err := c.Read()
	if err != nil {
		return nil, err
	}

	var buckets []*Bucket
	err = json.Unmarshal(bytes, &buckets)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func groupBuckets(buckets []*Bucket) map[string]*Bucket {
	bucketMap := make(map[string]*Bucket)
	for _, bucket := range buckets {
		bucketId := bucket.Id()
		bucketMap[bucketId] = bucket
	}
	return bucketMap
}

func flatBucketMap(bucketMap map[string]*Bucket) []*Bucket {
	buckets := make([]*Bucket, 0)
	for _, bucket := range bucketMap {
		buckets = append(buckets, bucket)
	}
	return buckets
}
