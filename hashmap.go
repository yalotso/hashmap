package main

import "sync"

const (
	loadFactor  = 0.75
	initialSize = 8
)

type HashMap struct {
	sync.RWMutex
	count      int
	leftBound  int
	rightBound int
	data       []*KeyValue
}

type KeyValue struct {
	key   interface{}
	value interface{}
}

var dummy = KeyValue{}

func NewHashMap() *HashMap {
	hm := HashMap{}
	hm.data = make([]*KeyValue, initialSize, initialSize)
	hm.rightBound = loadFactor * initialSize
	return &hm
}

func (hm *HashMap) Add(key int, value interface{}) {
	hm.Lock()
	isNew := addToData(hm.data, key, value)
	if isNew {
		hm.count++
	}
	if hm.count > hm.rightBound {
		hm.resize(cap(hm.data) << 1)
	}
	hm.Unlock()
}

func (hm *HashMap) Get(key int) interface{} {
	var result interface{}
	hm.RLock()
	i, ok := find(hm.data, key)
	if ok {
		result = hm.data[i].value
	}
	hm.RUnlock()
	return result
}

func (hm *HashMap) Delete(key int) bool {
	var deleted bool
	hm.Lock()
	i, ok := find(hm.data, key)
	if ok {
		hm.data[i] = &dummy
		hm.count--
		deleted = true
	}
	if hm.leftBound != 0 && hm.count < hm.leftBound {
		hm.resize(cap(hm.data) >> 1)
	}
	hm.Unlock()
	return deleted
}

func (hm *HashMap) resize(newsize int) {
	if cap(hm.data) < newsize {
		hm.leftBound, hm.rightBound = hm.rightBound, int(loadFactor*float64(newsize))
	} else if newsize == initialSize {
		hm.leftBound, hm.rightBound = 0, loadFactor*initialSize
	} else {
		hm.leftBound, hm.rightBound = int(loadFactor*float64(newsize>>1)), int(loadFactor*float64(newsize))
	}
	data := make([]*KeyValue, newsize, newsize)
	for _, item := range hm.data {
		if item == nil || item == &dummy {
			continue
		}
		key, ok := item.key.(int)
		if ok == false {
			continue
		}
		addToData(data, key, item.value)
	}
	hm.data = data
}

func addToData(data []*KeyValue, key int, value interface{}) bool {
	length := cap(data)
	for i := Index(key, length); ; i++ {
		if i == length {
			i = 0
		}
		if data[i] == nil || data[i] == &dummy {
			data[i] = &KeyValue{key, value}
			return true
		}
		if data[i].key == key {
			data[i].value = value
			return false
		}
	}
}

func find(data []*KeyValue, key int) (res int, found bool) {
	length := cap(data)
	start := Index(key, length)
	for i, ok := start, true; ok; i, ok = i+1, i != start {
		if i == length {
			i = 0
		}
		if data[i] == nil {
			break
		}
		if data[i].key == key {
			res, found = i, true
			break
		}
	}
	return
}
