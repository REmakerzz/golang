package main

import (
	"fmt"
	"testing"
)

func BenchmarkHashMapSet(b *testing.B) {
	m := NewHashMap(16, WithHashCRC32())
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key-%d", i), i)
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	m := NewHashMap(16, WithHashCRC32())
	for i := 0; i < b.N; i++ {
		m.Set(fmt.Sprintf("key-%d", i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(fmt.Sprintf("key-%d", i))
	}
}
