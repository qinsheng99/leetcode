package main

import (
	"container/list"
)

type LRUCache struct {
	maxBytes int
	l        *list.List
	cache    map[int]*list.Element
}

type valueCache struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxBytes: capacity,
		l:        list.New(),
		cache:    make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.l.MoveToBack(v)

	return v.Value.(*valueCache).val
}

func (this *LRUCache) Put(key int, value int) {
	vv, ok := this.cache[key]
	if ok {
		vv.Value.(*valueCache).val = value

		this.l.MoveToBack(vv)
		return
	}

	if this.l.Len() >= this.maxBytes {
		delete(this.cache, this.l.Front().Value.(*valueCache).key)
		this.l.Remove(this.l.Front())
	}

	this.cache[key] = this.l.PushBack(&valueCache{key: key, val: value})
}
