package dupcheck

func NewDupeCheck(size int) DupeCheck {
	if size <= 100 {
		return NewSliceBackedDupeCheck(size)
	}
	return NewMapBackedDupeCheck(size)
}
