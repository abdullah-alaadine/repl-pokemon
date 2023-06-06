package pokecache

import "testing"

func TestCreateTest(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Errorf("key1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("value doesn't match")
	}
}
