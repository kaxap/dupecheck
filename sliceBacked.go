package dupcheck

import "reflect"

type SliceBacked struct {
	Ring         []Item
	currentIndex int
	overflowFlag bool
}

func (s *SliceBacked) Add(item Item) Item {
	if item == nil || reflect.ValueOf(item).IsNil() {
		panic("Cannot add nil value to SliceBacked")
	}

	i := s.find(item)
	if i > 0 {
		return item
	}

	var old Item = nil
	if s.overflowFlag {
		old = s.Ring[s.currentIndex]
	}
	s.Ring[s.currentIndex] = item
	if s.currentIndex == len(s.Ring)-1 {
		s.currentIndex = 0
		s.overflowFlag = true
	} else {
		s.currentIndex++
	}
	return old
}

func (s *SliceBacked) Delete(item Item) bool {
	if item == nil || reflect.ValueOf(item).IsNil() {
		panic("Cannot delete nil value from SliceBacked")
	}

	i := s.find(item)
	if i < 0 {
		return false
	}

	s.Ring = append(s.Ring[:i], s.Ring[i+1:len(s.Ring)]...)
	s.Ring = append(s.Ring, nil)
	return true
}

func (s *SliceBacked) find(item Item) int {
	if item == nil || reflect.ValueOf(item).IsNil() {
		return -1
	}

	if s.overflowFlag {
		for i := 0; i < len(s.Ring); i++ {
			if item.Equal(s.Ring[i]) {
				return i
			}
		}
		return -1
	} else {
		for i := 0; i < s.currentIndex; i++ {
			if item.Equal(s.Ring[i]) {
				return i
			}
		}
		return -1
	}
}

func (s *SliceBacked) Has(item Item) bool {
	return s.find(item) >= 0
}

func NewSliceBacked(size int) DupeCheck {
	return &SliceBacked{Ring: make([]Item, size, size), currentIndex: 0, overflowFlag: false}
}
