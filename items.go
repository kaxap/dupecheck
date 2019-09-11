package dupcheck

import "reflect"

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

type StringItem struct {
	i string
}

func (e *StringItem) Value() interface{} {
	return e.i
}

func (e *StringItem) Equal(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	return e.i == item.(*StringItem).i
}
