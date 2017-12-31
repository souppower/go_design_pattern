package iterator

// Aggregate is a collection interface
type Aggregate interface {
	Iterator() Iterator
}

// Iterator is an iterator interface
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// BookShelf implements Aggregate
type BookShelf struct {
	Books []*Book
}

// Add adds a book to the bookshelf
func (bs *BookShelf) Add(book *Book) {
	bs.Books = append(bs.Books, book)
}

// Iterator returns an Iterator instance
func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: bs}
}

// BookShelfIterator implements Iterator
type BookShelfIterator struct {
	*BookShelf
	last int
}

// HasNext tells if the iterator can return an item
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.last >= len(bsi.Books) {
		return false
	}
	return true
}

// Next returns the current item, and advances the iterator
func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.Books[bsi.last]
	bsi.last++
	return book
}

// Book represents a struct stored in the BookShelf
type Book struct {
	name string
}
