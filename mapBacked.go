package dupecheck

import "reflect"

type MapBacked struct {
	sliceCheck DupeCheck
	m          map[interface{}]struct{}
}

func (m *MapBacked) Add(item Item) Item {
	old := m.sliceCheck.Add(item)
	if old != nil {
		// evict old item
		delete(m.m, old.Value())
	}

	// add new
	m.m[item.Value()] = struct{}{}
	return old
}

func (m *MapBacked) Delete(item Item) bool {
	if m.sliceCheck.Delete(item) {
		delete(m.m, item.Value())
		return true
	}
	return false
}

func (m *MapBacked) Has(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	_, ok := m.m[item.Value()]
	return ok
}

func NewMapBacked(size int) DupeCheck {
	return &MapBacked{sliceCheck: NewSliceBacked(size), m: make(map[interface{}]struct{})}
}
