package dupcheck

type Item interface {
	Equal(item Item) bool
	Value() interface{}
}
