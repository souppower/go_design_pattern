package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	p1 := Player{Name: "A", Strategy: &winningStrategy{seed: 10}}
	p2 := Player{Name: "B", Strategy: &winningStrategy{seed: 20}}

	p1h := p1.NextHand()
	p2h := p2.NextHand()

	if p1h.IsStrongerThan(p2h) {
		p1.Win()
		p2.Lose()
	} else if p1h.IsWeakerThan(p2h) {
		p1.Lose()
		p2.Win()
	} else {
		p1.Even()
		p2.Even()
	}
}
