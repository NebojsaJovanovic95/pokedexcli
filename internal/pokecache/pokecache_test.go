package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected value %q, got %q", c.val, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const shortInterval = 10 * time.Millisecond
	const waitTime = 20 * time.Millisecond

	cache := NewCache(shortInterval)
	cache.Add("https://example.com", []byte("testdata"))

	// Ensure the key exists initially
	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key initially")
		return
	}

	// Wait for longer than the interval
	time.Sleep(waitTime)

	// Now the key should be removed by the reaper
	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected key to be removed after interval")
		return
	}
}
