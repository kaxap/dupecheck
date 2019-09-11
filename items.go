package dupecheck

import "reflect"

type IntItem struct {
	V int
}

func (e *IntItem) Value() interface{} {
	return e.V
}

func (e *IntItem) Equal(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	return e.V == item.(*IntItem).V
}

type StringItem struct {
	V string
}

func (e *StringItem) Value() interface{} {
	return e.V
}

func (e *StringItem) Equal(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	return e.V == item.(*StringItem).V
}
