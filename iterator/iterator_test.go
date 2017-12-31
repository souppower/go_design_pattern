package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := &BookShelf{}

	asserts := []string{"A", "B", "C"}

	for _, assert := range asserts {
		bookShelf.Add(&Book{assert})
	}

	var i int
	bsi := bookShelf.Iterator()
	for bsi.HasNext() {
		if book := bsi.Next(); book.(*Book).name != asserts[i] {
			t.Errorf("Expected Book.name to be %s, but got %s", asserts[i], book.(*Book).name)
		}
		i++
	}

}
