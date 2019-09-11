package dupcheck

import (
	"testing"
)

func TestNewDupeCheck(t *testing.T) {
	d := New(100)
	switch d.(type) {
	case *SliceBacked:
		// ok
	default:
		t.Error("New has failed to create SliceBacked for size=100")
	}

	d = New(101)
	switch d.(type) {
	case *MapBacked:
		// ok
	default:
		t.Error("New has failed to create MapBacked for size=101")
	}

}
