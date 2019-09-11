package dupcheck

import (
	"testing"
)

func TestNewDupeCheck(t *testing.T) {
	d := New(100)
	switch d.(type) {
	case *SliceBackedDupeCheck:
		// ok
	default:
		t.Error("New has failed to create SliceBackedDupeCheck for size=100")
	}

	d = New(101)
	switch d.(type) {
	case *MapBackedDupeCheck:
		// ok
	default:
		t.Error("New has failed to create MapBackedDupeCheck for size=101")
	}

}
