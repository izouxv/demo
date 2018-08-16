package main

import (
	"fmt"
	"mynotes/cache"
	"strconv"
)

func main() {
	cacheA := cache.Cache("myCache")

	// The data loader gets called automatically whenever something
	// tries to retrieve a non-existing key from the cache.
	cacheA.SetDataLoader(func(key interface{}, args ...interface{}) *cache.CacheItem {
		// Apply some clever loading logic here, e.g. read values for
		// this key from database, network or file.
		val := "This is a test with key " + key.(string)

		// This helper method creates the cached item for us. Yay!
		item := cache.NewCacheItem(key, 0, val)
		return item
	})

	// Let's retrieve a few auto-generated items from the cache.
	for i := 0; i < 10; i++ {
		res, err := cacheA.Value("someKey_" + strconv.Itoa(i))
		if err == nil {
			fmt.Println("Found value in cache:", res.Data())
		} else {
			fmt.Println("Error retrieving value from cache:", err)
		}
	}
}
