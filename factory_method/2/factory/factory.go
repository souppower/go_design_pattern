package factory

// Product interface
type Product interface {
	GetOwner() string
}

type creator interface {
	New(owner string) Product
}

// Factory is a struct for the Create method
type Factory struct {
}

// Create returns a new instance using the factory that was passed in
func (f *Factory) Create(factory creator, owner string) Product {
	return factory.New(owner)
}
