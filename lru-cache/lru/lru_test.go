package lru

import "testing"

func TestLRUCache_BasicOperations(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put(1, 100)
	cache.Put(2, 200)

	if val := cache.Get(1); val != 100 {
		t.Errorf("Expected Get(1) to be 100, got %d", val)
	}

	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected Get(2) to be 200, got %d", val)
	}
}

func TestLRUCache_EvictionPolicy(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300) // Evicts key 1

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected key 1 to be evicted, but got %d", val)
	}

	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected Get(2) to be 200, got %d", val)
	}

	if val := cache.Get(3); val != 300 {
		t.Errorf("Expected Get(3) to be 300, got %d", val)
	}
}

func TestLRUCache_UpdateKey(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(1, 101) // Updates key 1

	if val := cache.Get(1); val != 101 {
		t.Errorf("Expected updated value for key 1 to be 101, got %d", val)
	}

	cache.Put(3, 300) // Evicts key 2

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected key 2 to be evicted, but got %d", val)
	}
}

func TestLRUCache_DeleteOperation(t *testing.T) {
	cache := NewLruCache(2)

	cache.Put(1, 100)
	cache.Put(2, 200)

	cache.Delete(1)
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected key 1 to be deleted, but got %d", val)
	}

	cache.Put(3, 300)

	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected Get(2) to still be 200, got %d", val)
	}
}

func TestLRUCache_CapacityOne(t *testing.T) {
	cache := NewLruCache(1)

	cache.Put(1, 100)
	if val := cache.Get(1); val != 100 {
		t.Errorf("Expected Get(1) to be 100, got %d", val)
	}

	cache.Put(2, 200) // Evicts key 1
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected key 1 to be evicted, but got %d", val)
	}
	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected Get(2) to be 200, got %d", val)
	}
}

func TestLRUCache_EdgeCases(t *testing.T) {
	cache := NewLruCache(2)

	if val := cache.Get(99); val != -1 {
		t.Errorf("Expected Get(99) to be -1, got %d", val)
	}

	cache.Delete(99) // Should handle gracefully without panic
}
