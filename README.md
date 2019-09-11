[![Build Status](https://travis-ci.org/kaxap/dupecheck.svg?branch=master)](https://travis-ci.org/kaxap/dupecheck)
[![Coverage Status](https://coveralls.io/repos/github/kaxap/dupecheck/badge.svg?branch=master)](https://coveralls.io/github/kaxap/dupecheck?branch=master)

# Description

A very simple library for checking some data for duplicates which is useful for a long-running applications checking for duplicates against values with potentially high cardinality.

For example, in an ETL application where the names of ingested files are checked against a list of previously ingested file names. In this case, maintaining a simple map can potentially lead to memory issues. 

It maintains a buffer of an arbitrary fixed size to and checks given items against it.
The `New(size int)` creates a slice-backed dupe checker for size <= 100, otherwise a map-backed dupe checker will be created.
The `NewSliceBacked(size int)` and `NewMapBacked(size int)` methods can be used directly.  

The interfaces:

```go

type DupeCheck interface {
  Add(item Item) Item
  Delete(item Item) bool
  Has(item Item) bool
}
```
and
```go

type Item interface {
  Equal(item Item) bool
  Value() interface{}
}

```
# Examples

An item with int value:
```go

type IntItem struct {
	i int
}

func (e *IntItem) Value() interface{} {
	return e.i
}

func (e *IntItem) Equal(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	return e.i == item.(*IntItem).i
}
```

Duplicate checker:

```go
package main

import (
	"fmt"
	"github.com/kaxap/dupecheck"
)

func main() {
	check := dupecheck.New(10)
	for i := 0; i < 10; i++ {
		_ = check.Add(&dupecheck.IntItem{i})
	}

	for i := 0; i < 20; i++ {
		if check.Has(&dupecheck.IntItem{i}) {
			fmt.Printf("%d is duplicated\n", i)
		} else {
			fmt.Printf("%d is not duplicated\n", i)
		}
	}
}
```

Output:
```
0 is duplicated
1 is duplicated
2 is duplicated
3 is duplicated
4 is duplicated
5 is duplicated
6 is duplicated
7 is duplicated
8 is duplicated
9 is duplicated
10 is not duplicated
11 is not duplicated
12 is not duplicated
13 is not duplicated
14 is not duplicated
15 is not duplicated
16 is not duplicated
17 is not duplicated
18 is not duplicated
19 is not duplicated
```
