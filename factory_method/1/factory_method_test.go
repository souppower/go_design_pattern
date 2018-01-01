package factoryMethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	f := &IDCardFactory{}
	products := []User{
		f.Create(f, "A"),
		f.Create(f, "B"),
		f.Create(f, "C"),
	}

	for i, product := range products {
		if owner := product.Use(); owner != assert[i] {
			t.Errorf("Expected owner to be %s, but got %s.\n", assert[i], owner)
		}
	}

}
