package dupecheck

func New(size int) DupeCheck {
	if size <= 100 {
		return NewSliceBacked(size)
	}
	return NewMapBacked(size)
}
