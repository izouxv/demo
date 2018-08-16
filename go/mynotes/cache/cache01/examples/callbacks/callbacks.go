package main

import (
	"fmt"
	"time"

	"mynotes/cache"
)

func main() {
	cacheA := cache.Cache("myCache")

	// This callback will be triggered every time a new item
	// gets added to the cache.
	cacheA.SetAddedItemCallback(func(entry *cache.CacheItem) {
		fmt.Println("Added:", entry.Key(), entry.Data(), entry.CreatedOn())
	})
	// This callback will be triggered every time an item
	// is about to be removed from the cache.
	cacheA.SetAboutToDeleteItemCallback(func(entry *cache.CacheItem) {
		fmt.Println("Deleting:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	// Caching a new item will execute the AddedItem callback.
	cacheA.Add("someKey", 0, "This is a test!")

	// Let's retrieve the item from the cache
	res, err := cacheA.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data())
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Deleting the item will execute the AboutToDeleteItem callback.
	cacheA.Delete("someKey")

	// Caching a new item that expires in 3 seconds
	res = cacheA.Add("anotherKey", 3*time.Second, "This is another test")

	// This callback will be triggered when the item is about to expire
	res.SetAboutToExpireCallback(func(key interface{}) {
		fmt.Println("About to expire:", key.(string))
	})

	time.Sleep(5 * time.Second)
}
