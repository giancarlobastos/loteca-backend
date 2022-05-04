package service

import (
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService() *CacheService {
	ristrettoCache, err := ristretto.NewCache(
		&ristretto.Config{
			NumCounters: 200000000,
			MaxCost:     20000000,
			BufferItems: 64,
		})

	if err != nil {
		log.Printf("Error [NewCacheService]: %v", err)
		return nil
	}

	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	cache := cache.New(ristrettoStore)

	return &CacheService{
		cache: cache,
	}
}

func (cs *CacheService) Get(key string) (interface{}, error) {
	value, err := cs.cache.Get(key)
	
	if err != nil {
		log.Printf("Error [Get]: %v - [%v]", err, key)
		return nil, err
	}
	
	return value, nil
}

func (cs *CacheService) Put(key string, value interface{}) error {

	if value == nil {
		return nil
	}

	return cs.cache.Set(key, value, &store.Options{Expiration: time.Minute * 30})
}

func (cs *CacheService) Delete(key string) error {
	return cs.cache.Delete(key)
}
