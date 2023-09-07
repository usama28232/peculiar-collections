// This package contains implmentations of Ordered Map & Slice/List with some general redundant operations
package peculiar

import "fmt"

// peculiar-map: a collection implementation for maintaining insertion order.
type Map[K comparable, T any] struct {
	keys *List[K]
	data map[K]T
}

// Creates new key value collection of provided type.
//
// Returns new instance of peculiar-map
func NewMap[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		keys: NewList[K](),
		data: make(map[K]T),
	}
}

// Creates new key value collection of provided type with given size.
//
// Returns new instance of Map
func NewMapOfSize[K comparable, T any](size int) *Map[K, T] {
	return &Map[K, T]{
		keys: NewList[K](),
		data: make(map[K]T, size),
	}
}

// Creates new key value collection of provided map.
//
// Returns new instance of Map
func NewMapWith[K comparable, T any](m map[K]T) *Map[K, T] {
	return &Map[K, T]{
		keys: NewList[K](),
		data: m,
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
		return value, fmt.Errorf(_KEY_NOT_EXISTS_ERR, key)
	}
}

// Checks if value exsists in collection.
//
// Returns true if value exists.
func (c *Map[K, T]) ContainsKey(key K) bool {
	return c.keys.Contains(key)
}

// Sets/Replaces value in collection.
//
// Return void
func (c *Map[K, T]) Set(key K, value T) {
	if _, ok := c.data[key]; !ok {
		c.keys.AddIfAbsent(key)
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
		err := c.keys.RemoveByValue(key)
		if err != nil {
			delete(c.data, key)
		}
	}
}

// Iterates current collection values by insertion order.
//
// Return void
func (c *Map[K, T]) Foreach(f func(v T)) {
	c.keys.Foreach(func(i int, v K) {
		value, _ := c.Get(v)
		f(value)
	})
}

// Iterates current collection with key value pair by insertion order.
//
// Return void
func (c *Map[K, T]) Entry(f func(k K, v T)) {
	c.keys.Foreach(func(i int, k K) {
		value, _ := c.Get(k)
		f(k, value)
	})
}

// Iterates current collection by insertion order and modify values foreach entry
//
// Return void
func (c *Map[K, T]) Map(f func(v T) T) {
	c.keys.Foreach(func(i int, k K) {
		value, _ := c.Get(k)
		newValue := f(value)
		c.Set(k, newValue)
	})
}

// Gets all keys in collection.
//
// Returns keys by insertion order.
func (c *Map[K, T]) Keys() []K {
	return c.keys.GetValues()
}

// Gets collection size
//
// Returns current collection size
func (c *Map[K, T]) Size() int {
	return c.keys.Length()
}

// Checks if collection is empty
//
// Returns true if collection is empty
func (c *Map[K, T]) IsEmpty() bool {
	return c.keys.IsEmpty()
}

// Clears all values in collection
//
// Return void
func (c *Map[K, T]) Clear() {
	if !c.IsEmpty() {
		c.keys.Clear()
		c.data = make(map[K]T)
	}
}

// Concats provided peculiar-map.
//
// NOTE: Provided Map will NOT OVERWRITE the existing Map values in case of conflict
//
// Returns new instance of peculiar-map
func (c *Map[K, T]) ConcatWithPrePrecedence(_map *Map[K, T]) *Map[K, T] {
	newMap := NewMapWith[K, T](c.data)
	_map.Entry(func(k K, v T) {
		newMap.SetIfAbsent(k, v)
	})
	return newMap
}

// Concats provided peculiar-map.
//
// NOTE: Provided Map will OVERWRITE the existing Map values in case of conflict
//
// Returns new instance of peculiar-map
func (c *Map[K, T]) ConcatWithPostPrecedence(_map *Map[K, T]) *Map[K, T] {
	newMap := NewMapWith[K, T](c.data)
	_map.Entry(func(k K, v T) {
		newMap.Set(k, v)
	})
	return newMap
}
