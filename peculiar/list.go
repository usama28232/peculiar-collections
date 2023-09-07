package peculiar

import (
	"errors"
)

// peculiar-list: a linear collection providing general operations over list/slice
type List[K comparable] struct {
	fake K
	data []K
}

// Creates new List/Slice of provided type
//
// Returns new instance of peculiar-list
func NewList[K comparable]() *List[K] {
	return &List[K]{
		data: []K{},
	}
}

// Creates new List/Slice of provided type
//
// NOTE: This will not remove duplicates from peculiar-list, to remove them see: `peculiar.List.RemoveDuplicates`
//
// Returns new instance of peculiar-list
func NewListWith[K comparable](k ...K) *List[K] {
	lst := NewList[K]()
	lst.data = append(lst.data, k...)
	return lst
}

// Gets length of collection
//
// Returns length of current collection
func (d *List[K]) Length() int {
	return len(d.data)
}

// Clears all values in collection
//
// Return void
func (d *List[K]) Clear() {
	d.data = []K{}
}

// Adds value in collection.
//
// NOTE: This will add new value in peculiar-list, and will not care about duplicates. To remove them see: `peculiar.List.RemoveDuplicates`
//
// To Update values, see: `peculiar.List.SetValue`
//
// Return void
func (d *List[K]) Add(item K) {
	d.data = append(d.data, item)
}

// Adds value in collection (only if absent).
//
// Return void
func (d *List[K]) AddIfAbsent(item K) {
	exists := d.Contains(item)
	if !exists {
		d.Add(item)
	}
}

// Removes value from collection on specified index.
//
// Return void
func (d *List[K]) Remove(index int) error {
	if index >= 0 && index < d.Length() {
		d.data = append(d.data[:index], d.data[index+1:]...)
	}
	return errors.New(_INDEX_RANGE_ERR)
}

// Removes specified value from collection.
//
// Returns error if value not found
func (d *List[K]) RemoveByValue(value K) error {
	i, _, err := d.get(value)
	if err == nil {
		d.Remove(i)
		return nil
	}
	return err
}

// Gets value by index from collection.
//
// Returns value or error if index is out of range.
func (d *List[K]) GetValue(index int) (K, error) {
	if index >= 0 && index < d.Length() {
		return d.data[index], nil
	}
	return d.fake, errors.New(_INDEX_RANGE_ERR)
}

// Gets all values in collection.
//
// Returns values by insertion order.
func (c *List[K]) GetValues() []K {
	return c.data
}

// Updates value in collection at specified index.
//
// Return void
func (d *List[K]) SetValue(index int, value K) error {
	if index >= 0 && index < d.Length() {
		d.data[index] = value
		return nil
	}
	return errors.New(_INDEX_RANGE_ERR)
}

// Gets index, value of given comparable type
//
// Returns index, value or error if value not found
func (d *List[K]) get(item K) (int, K, error) {
	if !d.IsEmpty() {
		for i, k := range d.data {
			if k == item {
				return i, k, nil
			}
		}
		return -1, d.fake, errors.New(_NOT_FOUND_ERR)
	}
	return -1, d.fake, errors.New(_NOT_FOUND_ERR)
}

// Iterates current collection.
//
// Return void
func (c *List[K]) Foreach(f func(i int, v K)) {
	for i, k := range c.data {
		f(i, k)
	}
}

// Iterates current collection and modify values foreach entry.
//
// Return void
func (d *List[K]) Map(f func(v K) K) {
	for i, k := range d.data {
		newValue := f(k)
		d.SetValue(i, newValue)
	}
}

// Checks if collections contains a value.
//
// Returns true if list contains element.
func (d *List[K]) Contains(value K) bool {
	_, _, err := d.get(value)
	return err == nil
}

// Checks if collections is empty.
//
// Returns true for empty collection.
func (d *List[K]) IsEmpty() bool {
	return d.Length() == 0
}

// Filter current collection into a new collection. This will return a new empty collection if current collection is empty or filter conditions are not satisfied
//
// Returns new filtered peculiar-list instance
func (d *List[K]) Filter(f func(v K) bool) *List[K] {
	if !d.IsEmpty() {
		subList := NewList[K]()
		for _, k := range d.data {
			isTrue := f(k)
			if isTrue {
				subList.Add(k)
			}
		}
		return subList
	} else {
		return NewList[K]()
	}
}

// Removes duplicates from current collection (if any).
//
// Returns new instance of peculiar-list
func (d *List[K]) RemoveDuplicates() *List[K] {
	newList := NewList[K]()
	d.Foreach(func(i int, v K) {
		newList.AddIfAbsent(v)
	})
	return newList
}

// Concats provided peculiar-list with current peculiar-list
//
// NOTE: This will not remove duplicates from peculiar-list, to remove them see: `peculiar.List.RemoveDuplicates`
//
// Return new list or current list if provided peculiar-list is empty
func (d *List[K]) Concat(l *List[K]) *List[K] {
	if !l.IsEmpty() {
		return NewListWith[K](d.data...)
	}
	return d
}

// Clones the current state of peculiar-list
//
// Return new list or current list if provided peculiar-list is empty
func (d *List[K]) Clone() *List[K] {
	return NewListWith[K](d.data...)
}
