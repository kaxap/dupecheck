# Description

A very simple library for checking some data for duplicates.

It maintains a buffer of an arbitrary fixed size to and checks given items against it.


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

