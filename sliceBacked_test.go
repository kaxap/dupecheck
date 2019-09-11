package dupecheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceBackedDupeCheck_Add(t *testing.T) {
	check := NewSliceBacked(10)
	for i := 0; i < 10; i++ {
		old := check.Add(&IntItem{i})
		assert.Nil(t, old)
	}

	for i := 0; i < 20; i++ {
		old := check.Add(&IntItem{i})
		assert.NotNil(t, old)
	}

}

func TestSliceBackedDupeCheck_Add_NilValue(t *testing.T) {
	check := NewSliceBacked(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	check.Add(nil)
	assert.True(t, paniced)
}

func TestSliceBackedDupeCheck_Add_NilInterfaceValue(t *testing.T) {
	check := NewSliceBacked(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	var item *IntItem = nil
	var item2 Item = item
	check.Add(item2)
	assert.True(t, paniced)
}

func TestSliceBackedDupeCheck_Delete_NilValue(t *testing.T) {
	check := NewSliceBacked(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	check.Delete(nil)
	assert.True(t, paniced)
}

func TestSliceBackedDupeCheck_Delete_NilInterfaceValue(t *testing.T) {
	check := NewSliceBacked(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	var item *IntItem = nil
	var item2 Item = item
	check.Delete(item2)
	assert.True(t, paniced)
}

func TestSliceBackedDupeCheck_Has_NilValue(t *testing.T) {
	check := NewSliceBacked(10)
	assert.False(t, check.Has(nil))
}

func TestSliceBackedDupeCheck_Has_NilInterfaceValue(t *testing.T) {
	check := NewSliceBacked(10)
	var item *IntItem = nil
	var item2 Item = item
	assert.False(t, check.Has(item2))
}

func TestSliceBackedDupeCheck_Delete(t *testing.T) {
	check := NewSliceBacked(10)
	for i := 0; i < 10; i++ {
		_ = check.Add(&IntItem{i})
	}

	assert.True(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{10}))
	for i := 1; i < 10; i++ {
		assert.True(t, check.Has(&IntItem{i}))
	}

	check = NewSliceBacked(10)
	for i := 0; i < 5; i++ {
		_ = check.Add(&IntItem{i})
	}

	assert.True(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{10}))
}

func TestSliceBackedDupeCheck_Has(t *testing.T) {
	check := NewSliceBacked(10)
	for i := 0; i < 10; i++ {
		_ = check.Add(&IntItem{i})
	}
	assert.True(t, check.Has(&IntItem{0}))
	assert.False(t, check.Has(&IntItem{10}))

	check = NewSliceBacked(10)
	for i := 0; i < 5; i++ {
		_ = check.Add(&IntItem{i})
	}
	assert.True(t, check.Has(&IntItem{0}))
	assert.False(t, check.Has(&IntItem{10}))
}
