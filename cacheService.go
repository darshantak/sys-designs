package main

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func newLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}
func (lru *LRUCache) get(key int) (int, bool) {
	if elem, exists := lru.cache[key]; exists {
		lru.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return -1, false
}
func (lru *LRUCache) put(key, value int) {
	if elem, exists := lru.cache[key]; exists {
		elem.Value.(*entry).value = value
		lru.list.MoveToFront(elem)
		return
	}
	if lru.list.Len() == lru.capacity {
		lru.evict()
	}
	newEntry := &entry{key: key, value: value}
	elem := lru.list.PushFront(newEntry)
	lru.cache[key] = elem
}
func (lru *LRUCache) evict() {
	lruElem := lru.list.Back()
	if lruElem != nil {
		lru.list.Remove(lruElem)
		delete(lru.cache, lruElem.Value.(*entry).key)
	}
}
