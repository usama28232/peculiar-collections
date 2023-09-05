# peculiar-map

Custom Ordered Map Implementation in **Golang**

#### NOTE: This repository requires GO Version >= 1.18


A map implementation that preserves insertion order while using struct and also offers some other functions like:

| **Function** | **Description**|
|:-: |:-: |
| `Set` | Adds value to collection |
| `SetIfAbsent` | Adds value to collection only if not exists |
| `Get` | Gets Value from collection |
| `ContainsKey` | Checks if "key" is present in the collection |
| `IsEmpty` | Checks if collection is empty |
| `Size` | Returns current size of collection |
| `Foreach` | Executes given function over a collection (in insertion order) without modifications |
| `Map` | Executes given function over a collection (in insertion order) with modifications |
| `Keys` | Returns current keys by insertion order |
| `Clear` | Clears all contents in the collection |


## Installation

Import the package using the go get command

```
go get github.com/usama28232/peculiar-map
```

## Usage

To create a new key value pair instance for int

```
collection := peculiar.NewMap[int, SampleStruct]()
```

OR 

```
peculiar.NewMap[string, SampleStruct]()
```

You can also initialise map with specific size

```
peculiar.NewMapOfSize[string, SampleStruct](4)
```

Where `SampleStruct` is 

```
type SampleStruct struct {
	FirstName string
	LastName  string
	Age       int
}
```

## Insert Values 

To add values to collection

```
...
collection := peculiar.NewMap[int, SampleStruct]()

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
...
```

You can use the same function to update existing values without disturbing the insertion order

```
...
sample, _ := collection.Get(2)
sample.Age++
collection.Set(2, sample)
...
```

## Get Values

To get a value from collection 

**NOTE:** Get function returns a value and an error

```
...
v, e := collection.Get(2)
if e != nil {
	// error case
	...
} else {
	fmt.Println(v)
}
...
```

## Remove Values

To remove the values from collection

```
...
collection.Remove(2)
...
```

## Miscellaneous Functions

You can iterate over collection using `Foreach` function

```
...
var sumOfAges = 0
collection.Foreach(func(v SampleStruct) {
	sumOfAges += v.Age
	fmt.Println(v)
})
fmt.Println("Sum of Ages:", sumOfAges)
...
```

Or you may use `Map` function if you want to modify underlying values

```
...
collection.Map(func(v SampleStruct) SampleStruct {
	v.Age++
	fmt.Println(v)
	return v
})
...
```

**Note:** *Above function will modify the underlying collection, proceed with caution!*

If you wish to get all keys from the collection, you may use

```
...
collection.Keys()
...
```

... Hope that helps


## Problem

In GO-lang, a map of struct does not guarantees insertion order while iterating over it, So I had to come up with a solution. 

While I was at it, I felt the need to use some generics to make it a more useable solution for better usage


### Feel free to edit/expand/explore this repository

For feedback and queries, reach me on LinkedIn [here](https://www.linkedin.com/in/usama28232/?original_referer=)
