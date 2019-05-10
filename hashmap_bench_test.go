package hashmap

import (
	"sync"
	"testing"
)

func BenchmarkHashMap_Add(b *testing.B) {
	m := NewHashMap()
	for i := 0; i < b.N; i++ {
		m.Add(i, i)
	}
}

func BenchmarkHashMap_Get(b *testing.B) {
	m := NewHashMap()
	for i := 0; i < b.N; i++ {
		_ = m.Get(i)
	}
}

func BenchmarkMap_Add(b *testing.B) {
	m := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkMap_Get(b *testing.B) {
	m := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func BenchmarkSyncMap_Add(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		m.Load(i)
	}
}
