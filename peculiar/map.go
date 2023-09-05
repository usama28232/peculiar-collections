package peculiar

import "fmt"

// peculiar-map - collection implementation for maintaining insertion order.
type Map[K comparable, T any] struct {
	keys []K
	data map[K]T
}

// Creates new key value collection of provided type.
//
// Returns new instance of peculiar-map
func NewMap[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		keys: []K{},
		data: make(map[K]T),
	}
}

// Creates new key value collection of provided type with given size.
//
// Returns new instance of Map
func NewMapOfSize[K int | string, T any](size int) *Map[K, T] {
	return &Map[K, T]{
		keys: []K{},
		data: make(map[K]T, size),
	}
}

// Gets value by key from collection.
//
// Returns value or error if key does not exists.
func (c *Map[K, T]) Get(key K) (T, error) {
	var value T
	if c.ContainsKey(key) {
		return c.data[key], nil
	} else {
		return value, fmt.Errorf("key '%v' does not exists", key)
	}
}

// Checks if value exsists in collection.
//
// Returns true if value exists.
func (c *Map[K, T]) ContainsKey(key K) bool {
	_, exists := c.data[key]
	return exists
}

// Sets/Replaces value in collection.
//
// Return void
func (c *Map[K, T]) Set(key K, value T) {
	if _, ok := c.data[key]; !ok {
		c.keys = append(c.keys, key)
	}
	c.data[key] = value
}

// Sets value in collection (only if absent).
//
// Return void
func (c *Map[K, T]) SetIfAbsent(key K, value T) {
	if !c.ContainsKey(key) {
		c.Set(key, value)
	}
}

// Removes key & value from collection.
//
// Return void
func (c *Map[K, T]) Remove(key K) {
	if c.ContainsKey(key) {
		delete(c.data, key)
		newKeys := []K{}
		for _, k := range c.keys {
			if k != key {
				newKeys = append(newKeys, k)
			}
		}
		c.keys = newKeys
	}
}

// Iterates current collection by insertion order.
//
// Return void
func (c *Map[K, T]) Foreach(f func(v T)) {
	for _, k := range c.keys {
		value, _ := c.Get(k)
		f(value)
	}
}

// Iterates current collection by insertion order and modify values foreach entry
//
// Return void
func (c *Map[K, T]) Map(f func(v T) T) {
	for _, k := range c.keys {
		value, _ := c.Get(k)
		newValue := f(value)
		c.Set(k, newValue)
	}
}

// Gets all keys in collection.
//
// Returns keys by insertion order.
func (c *Map[K, T]) Keys() []K {
	return c.keys
}

// Gets collection size
//
// Returns current collection size
func (c *Map[K, T]) Size() int {
	return len(c.keys)
}

// Checks if collection is empty
//
// Returns true if collection is empty
func (c *Map[K, T]) IsEmpty() bool {
	return len(c.keys) == 0
}

// Clears all values in collection
//
// Return void
func (c *Map[K, T]) Clear() {
	if !c.IsEmpty() {
		c.keys = []K{}
		c.data = make(map[K]T)
	}
}
