package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct {
	buckets    []map[string]interface{}
	hashFunc   func(key string) uint32
	bucketSize uint32
}

type HashOption func(*HashMap)

func NewHashMap(size int, options ...HashOption) *HashMap {
	h := &HashMap{
		buckets:    make([]map[string]interface{}, size),
		hashFunc:   defaultHash,
		bucketSize: uint32(size),
	}
	for i := range h.buckets {
		h.buckets[i] = make(map[string]interface{})
	}
	for _, option := range options {
		option(h)
	}
	return h
}

func (h *HashMap) Set(key string, value interface{}) {
	index := h.hashFunc(key) % h.bucketSize
	h.buckets[index][key] = value
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	index := h.hashFunc(key) % h.bucketSize
	value, ok := h.buckets[index][key]
	return value, ok
}

func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}

func defaultHash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func WithHashCRC64() HashOption {
	return func(h *HashMap) {
		table := crc64.MakeTable(crc64.ECMA)
		h.hashFunc = func(key string) uint32 {
			return uint32(crc64.Checksum([]byte(key), table))
		}
	}
}

func WithHashCRC32() HashOption {
	return func(h *HashMap) {
		h.hashFunc = defaultHash
	}
}

func WithHashCRC16() HashOption {
	return func(h *HashMap) {
		h.hashFunc = func(key string) uint32 {
			return uint32(len(key) % 65536)
		}
	}
}

func WithHashCRC8() HashOption {
	return func(h *HashMap) {
		h.hashFunc = func(key string) uint32 {
			return uint32(len(key) % 256)
		}
	}
}

func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})

	fmt.Println("CRC64:", since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})

	fmt.Println("CRC32:", since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println("CRC16:", since)

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println("CRC8:", since)
}
