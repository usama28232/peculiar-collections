# peculiar-collections 


[![GoDoc](https://godoc.org/pault.ag/go/piv?status.svg)](https://pkg.go.dev/github.com/usama28232/peculiar-collections) 

This repository is designed to centralize collection implementations with most commonly used functionalities to remove redundancy of code and provide fine grained control over general operations

**Peculiar-Collections v1.0.2** provides following implementations

* `Peculiar-Map` - is an implementation of Map which preserves insertion order while using struct and exposes general functions for re-usability
* `Peculiar-List` - is an implementation of slice/list which exposes general functions for re-usability

#### NOTE: This repository requires GO Version >= 1.18


## Installation

Import the package using the go get command

```
go get github.com/usama28232/peculiar-collections
```

### For more information & usage, please check out the wiki section

* [peculiar-map wiki](https://github.com/usama28232/peculiar-collections/wiki/peculiar%E2%80%90map)
* [peculiar-list wiki](https://github.com/usama28232/peculiar-collections/wiki/peculiar%E2%80%90list)


## Problem

In GO-lang, a map of struct does not guarantees insertion order while iterating over it, So I had to come up with a solution. 

While I was at it, I felt the need to use some generics to make it a more useable solution for better usage


### Feel free to edit/expand/explore this repository

For feedback and queries, reach me on LinkedIn [here](https://www.linkedin.com/in/usama28232/?original_referer=)
