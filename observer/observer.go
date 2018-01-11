package observer

import (
	"math/rand"
)

type numberGenerator struct {
	observers []observer
}

func (ng *numberGenerator) AddObserver(observer observer) {
	ng.observers = append(ng.observers, observer)
}

func (ng *numberGenerator) notifyObservers() []int {
	var result []int
	for _, observer := range ng.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}

type number interface {
	getNumber() int
}

func (ng *randomNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (ng *randomNumberGenerator) Execute() []int {
	return ng.notifyObservers()
}

type observer interface {
	update() int
}

type DigitObserver struct {
	generator number
}

func (do *DigitObserver) update() int {
	return do.generator.getNumber()
}
