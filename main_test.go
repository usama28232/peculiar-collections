package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/usama28232/peculiar-collections/peculiar"
)

type SampleStruct struct {
	FirstName string
	LastName  string
	Age       int
}

var collection *peculiar.Map[int, SampleStruct]
var list *peculiar.List[string]
var linkedList *peculiar.LinkedList[int]

func TestMain(m *testing.M) {
	collection = peculiar.NewMap[int, SampleStruct]()
	list = peculiar.NewList[string]()
	linkedList = peculiar.NewLinkedList[int]()

	code := m.Run()
	os.Exit(code)
}

func TestListInsertValue(t *testing.T) {
	list.Add("s0")
	list.Add("s1")
	list.Add("s2")
	list.Add("s3")
	list.Add("s4")
	list.Add("s5")
	list.Add("s5")
	printListValues(list)
}

func TestListRemoveValue(t *testing.T) {
	list.Remove(4)
	printListValues(list)
}

func TestListRemoveThrowsError(t *testing.T) {
	err := list.Remove(9)
	if err == nil {
		t.Fail()
	}
	printListValues(list)
}

func TestListGetValueByIndex(t *testing.T) {
	v, err := list.GetValue(1)
	if err != nil {
		t.Fail()
	}
	fmt.Println("Got value:", v)
}

func TestListGetValueByIndexWithIndexOutOfRange(t *testing.T) {
	v, err := list.GetValue(9)
	if err == nil {
		t.Fail()
	}
	fmt.Println("Got value:", v)
}

func TestListFilter(t *testing.T) {
	fmt.Println("List Filter test")
	value := "s5"
	newList := list.Filter(func(v string) bool {
		return v == value
	})
	if newList.Length() == 2 {
		printListValues(newList)
	} else {
		fmt.Println("filter not working", newList)
	}
}

func printListValues(lst *peculiar.List[string]) {
	fmt.Println("Printing List values")
	lst.Foreach(func(index int, value string) {
		fmt.Println(index, value)
	})
}

func TestMapInsertValue(t *testing.T) {
	sample1 := SampleStruct{
		FirstName: "John",
		LastName:  "Wick",
		Age:       33,
	}
	sample2 := SampleStruct{
		FirstName: "Chriss",
		LastName:  "Evans",
		Age:       32,
	}
	sample3 := SampleStruct{
		FirstName: "Johnny",
		LastName:  "Bravo",
		Age:       20,
	}
	collection.Set(1, sample1)
	collection.Set(2, sample2)
	collection.Set(3, sample3)
}

func TestMapInsertIfAbsebt(t *testing.T) {
	sample1 := SampleStruct{
		FirstName: "John",
		LastName:  "Wick",
		Age:       0,
	}
	collection.SetIfAbsent(1, sample1)

	v, _ := collection.Get(1)
	if v.Age == 0 {
		t.Fail()
	}
}

func TestMapGetValue(t *testing.T) {
	v, e := collection.Get(2)
	if e != nil {
		t.Fail()
	} else {
		fmt.Println(v)
	}
}

func TestMapUpdateValue(t *testing.T) {
	sample, _ := collection.Get(2)
	sample.Age++
	collection.Set(2, sample)
}

func TestMapValueNotExists(t *testing.T) {
	_, e := collection.Get(4)
	if e == nil {
		t.Fail()
	}
}

func TestMapIterate(t *testing.T) {
	collection.Map(func(v SampleStruct) SampleStruct {
		v.Age++
		fmt.Println(v)
		return v
	})
}

func TestMapIterateForeach(t *testing.T) {
	var sumOfAges = 0
	collection.Foreach(func(v SampleStruct) {
		sumOfAges += v.Age
		fmt.Println(v)
	})
	fmt.Println("Sum of Ages:", sumOfAges)
}

func TestMapRemoveValueByKey(t *testing.T) {
	collection.Remove(2)

	collection.Keys()

	collection.Foreach(func(v SampleStruct) {
		fmt.Println(v)
	})
}

func TestMapClear(t *testing.T) {
	collection.Clear()

	if !collection.IsEmpty() || collection.Size() > 0 {
		t.Fail()
	}
}

func TestLinkedListAddValue(t *testing.T) {
	linkedList.AddByValue(1)
	linkedList.AddByValue(2)
	linkedList.AddByValue(3)
	linkedList.AddByValue(4)
	linkedList.AddByValue(5)
	fmt.Println(linkedList.GetItemsSlice())
	if linkedList.Length() == 0 {
		t.Error("cannot add values to peculiar-linked-list")
	}
}

func TestLinkedListAddBeforeValue(t *testing.T) {
	itm2 := peculiar.NewLinkedListItemWithValue[int](2)
	itm4 := peculiar.NewLinkedListItemWithValue[int](4)
	err1 := linkedList.AddBeforeValue(3, itm2)
	err2 := linkedList.AddBeforeValue(4, itm4)
	fmt.Println(linkedList.GetItemsSlice())
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
	}
}

func TestLinkedListAddAfterValue(t *testing.T) {
	itm2 := peculiar.NewLinkedListItemWithValue[int](2)
	itm4 := peculiar.NewLinkedListItemWithValue[int](4)
	err1 := linkedList.AddAfterValue(3, itm2)
	err2 := linkedList.AddAfterValue(4, itm4)
	fmt.Println(linkedList.GetItemsSlice())
	if err1 != nil || err2 != nil {
		t.Error(err1, err2)
	}
}

func TestLinkedListRemoveValueAndFind(t *testing.T) {
	linkedList.RemoveByValue(3)
	fmt.Println(linkedList.GetItemsSlice())
	item, err := linkedList.Find(3)
	if err == nil {
		t.Error("could not delete value from peculiar-linked-list")
		fmt.Println(linkedList.GetItemsSlice(), item)
	}
}
