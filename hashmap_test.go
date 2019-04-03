package main

import (
	"sync"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	m := NewHashMap()
	if m.count != 0 {
		t.Error("map is not empty")
	}
	for _, v := range m.data {
		if v == nil {
			continue
		}
		if v.key != nil || v.value != nil {
			t.Error("map is not empty")
		}
	}
}

func TestAdd(t *testing.T) {
	m := NewHashMap()
	keys := []int{1, 2, 3, 4}
	value := "test value"
	for _, key := range keys {
		m.Add(key, value)
	}
	if m.count != cap(keys) {
		t.Errorf("map item count, expected: 4, actual %v", m.count)
	}
}

func TestGet(t *testing.T) {
	m := NewHashMap()
	v := m.Get(1)
	if v != nil {
		t.Error("item is not nil")
	}
	value := "test value"
	m.Add(1, value)
	v = m.Get(1)
	if v == nil {
		t.Error("item not found")
	}
	if v != value {
		t.Errorf("wrong value, expected: %v, actual: %v", value, v)
	}
}

func TestDelete(t *testing.T) {
	m := NewHashMap()
	m.Add(1, "test value")
	m.Delete(1)
	v := m.Get(1)
	if v != nil {
		t.Error("item is not nil")
	}
	if m.count != 0 {
		t.Error("map is not empty")
	}
}

func TestDelete_Size(t *testing.T) {
	m := NewHashMap()
	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, key := range keys {
		m.Add(key, key)
	}
	for i := len(keys) - 1; i >= 0; i-- {
		deleted := m.Delete(keys[i])
		if !deleted {
			t.Errorf("item with key %v has not been deleted", keys[i])
		}
		v := m.Get(keys[i])
		if v != nil {
			t.Error("item is not nil")
		}
	}
	if m.count != 0 {
		t.Error("map is not empty")
	}
	if cap(m.data) != 8 {
		t.Errorf("expected %v map size, actual %v", 8, cap(m.data))
	}
}

func TestOverwrite(t *testing.T) {
	m := NewHashMap()
	key1, value1 := 1, "test value 1"
	key2, value2 := 1, "test value 2"
	m.Add(key1, value1)
	m.Add(key2, value2)
	if m.count != 1 {
		t.Errorf("map item count, expected: 1, actual %v", m.count)
	}
	v := m.Get(1)
	if v != value2 {
		t.Error("wrong item")
	}
}

func TestResize(t *testing.T) {
	m := NewHashMap()
	oldsize := cap(m.data)
	m.resize(oldsize << 1)
	if cap(m.data) != oldsize*2 || cap(m.data) != oldsize*2 {
		t.Error("incorrect resize")
	}
}

func TestResizing(t *testing.T) {
	m := NewHashMap()
	itemCount := 48
	for i := 0; i < itemCount; i++ {
		m.Add(i, i)
	}
	if m.count != itemCount {
		t.Errorf("expected %v elements, actual %v", itemCount, m.count)
	}
	if cap(m.data) != 64 {
		t.Errorf("expected %v map size, actual %v", 64, cap(m.data))
	}
	m.Add(49, 49)
	if cap(m.data) != 128 {
		t.Errorf("expected %v map size, actual %v", 128, cap(m.data))
	}
}

func TestHashMap_ConcurrentOverwrite(t *testing.T) {
	m := NewHashMap()
	itemsCount := 1000
	var wg sync.WaitGroup
	for i := 0; i < itemsCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < itemsCount; i++ {
				m.Add(i, i)
			}
		}()
	}
	wg.Wait()
	if m.count != itemsCount {
		t.Errorf("expected %v elements, actual %v", itemsCount, m.count)
	}
}
