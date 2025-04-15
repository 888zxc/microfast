package service

import (
	"sync"
	"time"
)

// Cache 是一个简单的内存缓存实现
type Cache struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// NewCache 创建新的缓存
func NewCache() *Cache {
	cache := &Cache{
		data: make(map[string]cacheItem),
	}
	
	// 启动垃圾回收
	go cache.startGC()
	
	return cache
}

// Set 设置缓存项
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.data[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(duration),
	}
}

// Get 获取缓存项
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	item, found := c.data[key]
	if !found {
		return nil, false
	}
	
	if time.Now().After(item.expiration) {
		return nil, false
	}
	
	return item.value, true
}

// 垃圾回收
func (c *Cache) startGC() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, item := range c.data {
			if now.After(item.expiration) {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
}
