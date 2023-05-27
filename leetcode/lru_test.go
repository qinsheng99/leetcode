package main

import (
	"testing"
)

type cacheSlice []int

func (c cacheSlice) Len() uint {
	return uint(len(c))
}

func TestLRU(t *testing.T) {
	c := newCache(100)

	c.Add("a", cacheSlice([]int{1, 2, 3, 4}))
	c.Add("b", cacheSlice([]int{5, 6, 7, 8}))
	c.Add("b", cacheSlice([]int{9, 10, 11, 12}))

	c.Remove()
	if f := c.Front(); f != nil {
		t.Log(f.(cacheSlice))
	}

}
