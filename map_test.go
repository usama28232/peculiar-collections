package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/usama28232/peculiar-map/peculiar"
)

type SampleStruct struct {
	FirstName string
	LastName  string
	Age       int
}

var collection *peculiar.Map[int, SampleStruct]

func TestMain(m *testing.M) {
	collection = peculiar.NewMap[int, SampleStruct]()
	code := m.Run()
	os.Exit(code)
}

func TestInsertValue(t *testing.T) {
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
	// Create new instance of Peculiar Map
}

func TestGetValue(t *testing.T) {
	v, e := collection.Get(2)
	if e != nil {
		t.Fail()
	} else {
		fmt.Println(v)
	}
}

func TestUpdateValue(t *testing.T) {
	sample, _ := collection.Get(2)
	sample.Age++
	collection.Set(2, sample)
}

func TestValueNotExists(t *testing.T) {
	_, e := collection.Get(4)
	if e == nil {
		t.Fail()
	}
}

func TestIterateMap(t *testing.T) {
	collection.Map(func(v SampleStruct) SampleStruct {
		v.Age++
		fmt.Println(v)
		return v
	})
}

func TestIterateForeach(t *testing.T) {
	var sumOfAges = 0
	collection.Foreach(func(v SampleStruct) {
		sumOfAges += v.Age
		fmt.Println(v)
	})
	fmt.Println("Sum of Ages:", sumOfAges)
}

func TestRemoveValueByKey(t *testing.T) {
	collection.Remove(2)

	collection.Keys()

	collection.Foreach(func(v SampleStruct) {
		fmt.Println(v)
	})
}

func TestCollective(t *testing.T) {
	collection1 := peculiar.NewMap[int, SampleStruct]()
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
	collection1.Set(1, sample1)
	collection1.Set(2, sample2)
	collection1.Set(3, sample3)
	collection1.Remove(2)
}
