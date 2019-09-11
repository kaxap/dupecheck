package dupcheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntItem_Equal(t *testing.T) {
	a := IntItem{1}
	b := IntItem{2}
	c := IntItem{1}
	var d Item = nil
	assert.True(t, a.Equal(&c))
	assert.True(t, c.Equal(&a))
	assert.False(t, a.Equal(&b))
	assert.False(t, b.Equal(&a))
	assert.False(t, b.Equal(d))

}

func TestIntItem_Value(t *testing.T) {
	a := IntItem{1}
	assert.Equal(t, a.Value(), 1)
}

func TestStringItem_Equal(t *testing.T) {
	a := StringItem{"1"}
	b := StringItem{"2"}
	c := StringItem{"1"}
	var d Item = nil
	assert.True(t, a.Equal(&c))
	assert.True(t, c.Equal(&a))
	assert.False(t, a.Equal(&b))
	assert.False(t, b.Equal(&a))
	assert.False(t, b.Equal(d))

}

func TestStringItem_Value(t *testing.T) {
	a := StringItem{"1"}
	assert.Equal(t, a.Value(), "1")
}
