package lrucache

import (
	"container/list"
	"errors"
)

// LRUCache implements a non-thread safe fixed size LRU cache
type LRUCache struct {
	m map[interface{}]*list.Element
	l *list.List
	c int
}

type entry struct {
	key   interface{}
	value interface{}
}

// NewLRUCache construct an LRU of the given size
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity < 1 {
		return nil, errors.New("capacity must be greater than 0")
	}
	l := &LRUCache{
		m: make(map[interface{}]*list.Element),
		l: list.New(),
		c: capacity,
	}
	return l, nil
}

// Get looks up given key's value in cache
func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	ele, ok := l.m[key]
	if !ok {
		return nil, ok
	}
	l.l.MoveToFront(ele)
	return ele.Value.(*entry).value, true
}

// Add adds a value to the cache. Returns true if an eviction occurred.
func (l *LRUCache) Add(key, value interface{}) bool {
	ele, ok := l.m[key]
	// edit ele if the key exists
	if ok {
		l.l.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return false
	}
	// evict tail if cache is full
	if l.l.Len() == l.c {
		delete(l.m, l.l.Back().Value.(*entry).key)
		l.l.Remove(l.l.Back())
	}
	pair := &entry{key, value}
	ele = l.l.PushFront(pair)
	l.m[key] = ele
	return true
}

// Len returns current size of the cache.
func (l *LRUCache) Len() int {
	return l.l.Len()
}
