package dupcheck

import (
	"testing"
)

func TestNewDupeCheck(t *testing.T) {
	d := NewDupeCheck(100)
	switch d.(type) {
	case *SliceBackedDupeCheck:
		// ok
	default:
		t.Error("NewDupeCheck has failed to create SliceBackedDupeCheck for size=100")
	}

	d = NewDupeCheck(101)
	switch d.(type) {
	case *MapBackedDupeCheck:
		// ok
	default:
		t.Error("NewDupeCheck has failed to create MapBackedDupeCheck for size=101")
	}

}
