package main

import (
	"container/list"
	"sync"
)

type cache struct {
	maxLen  uint
	current uint
	l       *list.List
	cache   map[string]*list.Element
	mutex   sync.RWMutex
}

func newCache(max uint) *cache {
	return &cache{
		maxLen: max,
		l:      list.New(),
		cache:  make(map[string]*list.Element),
		mutex:  sync.RWMutex{},
	}
}

type entry struct {
	key   string
	value Len
}

type Len interface {
	Len() uint
}

func (c *cache) Remove() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	e := c.l.Back()

	if e != nil {
		v := e.Value.(*entry)
		delete(c.cache, v.key)
		c.l.Remove(e)

		c.current -= uint(len(v.key)) + v.value.Len()
	}
}

func (c *cache) Get(key string) (Len, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if e, ok := c.cache[key]; ok {
		v := e.Value.(*entry)
		c.l.MoveToBack(e)

		return v.value, ok
	}

	return nil, false
}

func (c *cache) Add(key string, value Len) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.maxLen < value.Len() {
		return
	}

	if e, ok := c.cache[key]; ok {
		v := e.Value.(*entry)
		c.current += value.Len() - v.value.Len()

		v.value = value

		c.l.MoveToBack(e)
	} else {
		front := c.l.PushFront(&entry{key: key, value: value})
		c.cache[key] = front
	}

	for c.maxLen != 0 && c.current > c.maxLen {
		c.Remove()
	}
}
