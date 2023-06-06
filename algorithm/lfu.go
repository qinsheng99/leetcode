package algorithm

import "container/list"

type lfuVC struct {
	value int
	count int
}

type lfuCache struct {
	cap int
	VC  map[int]lfuVC
	l   map[int]*list.List
	min int
}

func newLFU(cap int) *lfuCache {
	return &lfuCache{
		cap: cap,
		VC:  make(map[int]lfuVC, cap),
		l:   make(map[int]*list.List, cap),
	}
}

func (l *lfuCache) get(key int) int {
	v, ok := l.VC[key]
	if !ok {
		return -1
	}

	l.VC[key] = lfuVC{value: v.value, count: v.count + 1}

	for e := l.l[v.count].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			l.l[v.count].Remove(e)
			break
		}
	}

	if l.l[v.count+1] == nil {
		l.l[v.count+1] = list.New()
	}

	l.l[v.count+1].PushBack(key)

	if v.count == l.min && l.l[v.count].Len() == 0 {
		l.min++
	}

	return v.value
}

func (l *lfuCache) put(key, value int) {
	if _, ok := l.VC[key]; ok {
		l.VC[key] = lfuVC{value: value, count: l.VC[key].count}
		l.get(key)

		return
	}

	if len(l.VC) >= l.cap {
		ls := l.l[l.min]
		delete(l.VC, ls.Front().Value.(int))

		ls.Remove(ls.Front())
	}

	l.VC[key] = lfuVC{value: value, count: 1}

	if l.min != 1 {
		l.l[1] = list.New()
	}

	l.l[1].PushBack(key)

	l.min = 1
}
