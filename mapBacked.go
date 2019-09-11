package dupcheck

import "reflect"

type MapBackedDupeCheck struct {
	sliceCheck DupeCheck
	m          map[interface{}]struct{}
}

func (m *MapBackedDupeCheck) Add(item Item) Item {
	old := m.sliceCheck.Add(item)
	if old != nil {
		// evict old item
		delete(m.m, old.Value())
	}

	// add new
	m.m[item.Value()] = struct{}{}
	return old
}

func (m *MapBackedDupeCheck) Delete(item Item) bool {
	if m.sliceCheck.Delete(item) {
		delete(m.m, item.Value())
		return true
	}
	return false
}

func (m *MapBackedDupeCheck) Has(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return false
	}
	_, ok := m.m[item.Value()]
	return ok
}

func NewMapBackedDupeCheck(size int) DupeCheck {
	return &MapBackedDupeCheck{sliceCheck: NewSliceBackedDupeCheck(size), m: make(map[interface{}]struct{})}
}
