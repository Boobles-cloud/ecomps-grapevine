package utils

import (
	"sync"
	"time"

	"grapevine.boobles.cloud/global"
)

type CacheManager[T any] struct {
	items map[string]CacheEntry[T]
	lock  sync.Mutex
}

type CacheEntry[T any] struct {
	Item           T
	ExpirationTime time.Time
}

// Creates a new cache entry
func CreateNewCacheEntry[T any](item T) CacheEntry[T] {
	return CacheEntry[T]{
		Item:           item,
		ExpirationTime: time.Now().Add(2 * time.Hour),
	}
}

// Creates a new instance of cache manager
func CreateNewCacheManager[T any]() *CacheManager[T] {
	return &CacheManager[T]{
		items: make(map[string]CacheEntry[T]),
	}
}

// Sets or updates the wanted item
func (c *CacheManager[T]) SetOrUpdateItem(key string, item T) {

	c.lock.Lock()
	defer c.lock.Unlock()

	entry := CreateNewCacheEntry[T](item)
	c.items[key] = entry
}

// Returns a item by its key
func (c *CacheManager[T]) GetItem(key string) (T, error) {

	c.lock.Lock()
	defer c.lock.Unlock()

	var empty T

	if !c.keyExists(key) {
		return empty, global.ERR_KEY_NOT_FOUND
	}

	// Checks if the entry is expired
	if c.items[key].IsCacheEntryExpired() {
		c.RemoveItem(key)
		return empty, global.ERR_ENTRY_EXPIRED
	}

	return c.items[key].Item, nil
}

// Deprecated: DO NOT USE THIS!
// Gets all items
func (c *CacheManager[T]) GetItems(appId uint) ([]T, bool) {

	allItems := make([]T, 0, 100)

	// for k := range c.items {
	// 	if c.items[k].AppId == appId {
	// 		allItems = append(allItems, c.items[k].Item)
	// 	}
	// }

	return allItems, len(allItems) != 0
}

// Removes an item
func (c *CacheManager[T]) RemoveItem(key string) {

	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.keyExists(key) {
		return
	}

	delete(c.items, key)
}

// Checks if a key exists
func (c *CacheManager[T]) keyExists(key string) bool {
	_, ok := c.items[key]
	return ok
}

// Checks if the cache entry is expired
func (c CacheEntry[T]) IsCacheEntryExpired() bool {
	return time.Now().After(c.ExpirationTime)
}
