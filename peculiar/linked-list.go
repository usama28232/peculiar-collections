package peculiar

import (
	"errors"
	"sync"
)

// peculiar-linked-list: a linear collection providing general operations over linked-list
type LinkedList[K comparable] struct {
	size  int
	Items *LinkedListItem[K]
	last  *LinkedListItem[K]
	mu    sync.Mutex
}

// Creates new List/Slice of provided type
//
// Returns new instance of peculiar-linked-list
func NewLinkedList[K comparable]() *LinkedList[K] {
	return &LinkedList[K]{
		size: 0,
	}
}

// peculiar-linked-list-item: a collection item
type LinkedListItem[K comparable] struct {
	value K
	next  *LinkedListItem[K]
}

// Creates new List/Slice Node of provided type
//
// Returns new instance of peculiar-linked-list item
func NewLinkedListItem[K comparable]() *LinkedListItem[K] {
	return &LinkedListItem[K]{
		next: nil,
	}
}

// Creates new List/Slice Node of provided type
//
// Returns new instance of peculiar-linked-list item
func NewLinkedListItemWithValue[K comparable](value K) *LinkedListItem[K] {
	return &LinkedListItem[K]{
		value: value,
		next:  nil,
	}
}

// Gets peculiar-linked-list item value
//
// Return void
func (li *LinkedListItem[K]) GetValue() K {
	return li.value
}

// Sets peculiar-linked-list item value
//
// Return void
func (li *LinkedListItem[K]) SetValue(value K) {
	li.value = value
}

// Sets peculiar-linked-list next node
//
// Return void
func (li *LinkedListItem[K]) SetNext(value *LinkedListItem[K]) {
	li.next = value
}

// Gets peculiar-linked-list next node
//
// Return void
func (li *LinkedListItem[K]) GetNext() *LinkedListItem[K] {
	return li.next
}

// Checks if item has a next value
//
// Returns true if next value exists
func (li *LinkedListItem[K]) HasNext() bool {
	return li.next != nil
}

// Gets length of collection
//
// Returns length of current collection
func (l *LinkedList[K]) Length() int {
	return l.size
}

// Gets last item from collection
//
// Returns last item from collection
func (l *LinkedList[K]) getLastLink() *LinkedListItem[K] {
	lastLink := l.Items
	for {
		if lastLink.HasNext() {
			lastLink = lastLink.GetNext()
		} else {
			break
		}
	}
	return lastLink
}

// Adds item to Linked-List by value
//
// Return void
func (l *LinkedList[K]) AddByValue(value K) {
	item := NewLinkedListItem[K]()
	item.SetValue(value)
	l.Add(item)
}

// Adds item to Linked-List
//
// Return void
func (l *LinkedList[K]) Add(item *LinkedListItem[K]) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.size == 0 {
		l.Items = item
	} else {
		l.getLastLink().SetNext(item)
	}
	l.size++
	l.last = item
}

// Adds item to Linked-List at before first occurrence  of value
//
// Return void
func (l *LinkedList[K]) AddBeforeValue(value K, item *LinkedListItem[K]) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	valueFound := false
	ref := l.Items
	var prevValue *LinkedListItem[K]
	for {
		if value == ref.value {
			valueFound = true
			break
		}
		prevValue = ref
		ref = ref.GetNext()
		if ref == nil {
			break
		}
	}

	if valueFound {
		prevValue.SetNext(item)
		item.SetNext(ref)
		return nil
	}
	return errors.New(_NOT_FOUND_ERR)
}

// Adds item to Linked-List at after first occurrence of value
//
// Return void
func (l *LinkedList[K]) AddAfterValue(value K, item *LinkedListItem[K]) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	element, err := l.Find(value)
	if err != nil {
		return err
	}
	originalNext := element.GetNext()
	if originalNext == nil {
		element.SetNext(item)
	} else {
		element.SetNext(item)
		item.SetNext(originalNext)
	}
	return nil
}

// Removes by value from collection
//
// Return void
func (l *LinkedList[K]) RemoveByValue(value K) {
	l.mu.Lock()
	defer l.mu.Unlock()
	valueFound := false
	ref := l.Items
	var prevValue *LinkedListItem[K] = nil
	var nextValue *LinkedListItem[K] = nil
	for {
		if value == ref.value {
			valueFound = true
			nextValue = ref.GetNext()
			break
		}
		prevValue = ref
		ref = ref.next
		if ref == nil {
			break
		}
	}
	if valueFound {
		prevValue.SetNext(nextValue)
	}
}

// Finds an item in Linked-List
//
// Returns LinkedListItem or not found error
func (l *LinkedList[K]) Find(value K) (*LinkedListItem[K], error) {
	valueFound := false
	ref := l.Items
	for {
		if value == ref.value {
			valueFound = true
			break
		}
		ref = ref.GetNext()
		if ref == nil {
			break
		}
	}
	if valueFound {
		return ref, nil
	}
	return nil, errors.New(_NOT_FOUND_ERR)
}

// Gets first element from the linked list
//
// Returns Head of the Linked List or collection empty error
func (l *LinkedList[K]) Head() (*LinkedListItem[K], error) {
	if l.size > 0 {
		return l.Items, nil
	}
	return NewLinkedListItem[K](), errors.New(_COLLECTION_EMPTY)
}

// Gets last element from the linked list
//
// Returns Tail of the Linked List
func (l *LinkedList[K]) Tail(item LinkedListItem[K]) (*LinkedListItem[K], error) {
	if l.size > 0 {
		return l.last, nil
	}
	return NewLinkedListItem[K](), errors.New(_COLLECTION_EMPTY)
}

// Gets slice of items in peculiar-linked-list
//
// Returns items in a slice
func (l *LinkedList[K]) GetItemsSlice() []K {
	items := []K{}
	Item := l.Items
	for {
		items = append(items, Item.value)
		Item = Item.GetNext()
		if Item == nil {
			break
		}
	}
	return items
}
