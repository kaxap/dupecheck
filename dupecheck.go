package dupcheck

type DupeCheck interface {
	Add(item Item) Item
	Delete(item Item) bool
	Has(item Item) bool
}
