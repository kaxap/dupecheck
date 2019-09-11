package dupcheck

func New(size int) DupeCheck {
	if size <= 100 {
		return NewSliceBackedDupeCheck(size)
	}
	return NewMapBackedDupeCheck(size)
}
