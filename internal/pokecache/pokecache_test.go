package pokecache

import "testing"

func TestCreateTest(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}
