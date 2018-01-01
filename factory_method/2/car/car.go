package car

import (
	"../factory"
)

// Factory is for creating car instances
type Factory struct {
	*factory.Factory
}

// Car is a car product
type Car struct {
	owner string
}

// New returns a Car instance
func (cf *Factory) New(owner string) factory.Product {
	return &Car{owner}
}

// GetOwner returns the owner of its car
func (c *Car) GetOwner() string {
	return c.owner
}
