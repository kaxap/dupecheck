package dupcheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapBackedDupeCheck_Add(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	for i := 0; i < 10; i++ {
		old := check.Add(&IntItem{i})
		assert.Nil(t, old)
	}

	for i := 0; i < 20; i++ {
		old := check.Add(&IntItem{i})
		assert.NotNil(t, old)
	}

}

func TestMapBackedDupeCheck_Add_NilValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	check.Add(nil)
	assert.True(t, paniced)
}

func TestMapBackedDupeCheck_Add_NilInterfaceValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
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

func TestMapBackedDupeCheck_Delete_NilValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	paniced := false
	defer func() {
		_ = recover()
		paniced = false
	}()
	check.Delete(nil)
	assert.True(t, paniced)
}

func TestMapBackedDupeCheck_Delete_NilInterfaceValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
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

func TestMapBackedDupeCheck_Has_NilValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	assert.False(t, check.Has(nil))
}

func TestMapBackedDupeCheck_Has_NilInterfaceValue(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	var item *IntItem = nil
	var item2 Item = item
	assert.False(t, check.Has(item2))
}

func TestMapBackedDupeCheck_Delete(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	for i := 0; i < 10; i++ {
		_ = check.Add(&IntItem{i})
	}

	assert.True(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{10}))
	for i := 1; i < 10; i++ {
		assert.True(t, check.Has(&IntItem{i}))
	}

	check = NewMapBackedDupeCheck(10)
	for i := 0; i < 5; i++ {
		_ = check.Add(&IntItem{i})
	}

	assert.True(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{0}))
	assert.False(t, check.Delete(&IntItem{10}))
}

func TestMapBackedDupeCheck_Has(t *testing.T) {
	check := NewMapBackedDupeCheck(10)
	for i := 0; i < 10; i++ {
		_ = check.Add(&IntItem{i})
	}
	assert.True(t, check.Has(&IntItem{0}))
	assert.False(t, check.Has(&IntItem{10}))

	check = NewMapBackedDupeCheck(10)
	for i := 0; i < 5; i++ {
		_ = check.Add(&IntItem{i})
	}
	assert.True(t, check.Has(&IntItem{0}))
	assert.False(t, check.Has(&IntItem{10}))
}
