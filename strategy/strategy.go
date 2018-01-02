package strategy

import (
	"math/rand"
)

const (
	ROCK = iota
	SCISSOR
	PAPER
)

var hands []*hand

func init() {
	hands = []*hand{
		&hand{ROCK},
		&hand{SCISSOR},
		&hand{PAPER},
	}
}

type hand struct {
	handValue int
}

func getHand(handValue int) *hand {
	return hands[handValue]
}

func (sh *hand) IsStrongerThan(h *hand) bool {
	return sh.fight(h) == 1
}

func (sh *hand) IsWeakerThan(h *hand) bool {
	return sh.fight(h) == -1
}

func (sh *hand) fight(h *hand) int {
	if sh == h {
		return 0
	} else if (sh.handValue+1)%3 == h.handValue {
		return 1
	} else {
		return -1
	}
}

// Goでは関数がファーストクラスであるため、戦略がひとつのメソッドのみで
// 完結する場合は、関数を渡すことでStrategyパターンとすることもできる。
// type strategy func() *hand
type strategy interface {
	NextHand() *hand
	study(win bool)
}

type winningStrategy struct {
	won      bool
	prevHand *hand
}

func (ws *winningStrategy) NextHand() *hand {
	if !ws.won {
		ws.prevHand = getHand(rand.Intn(3))
	}
	return ws.prevHand
}

func (ws *winningStrategy) study(win bool) {
	ws.won = win
}

type Player struct {
	Name                           string
	Strategy                       strategy
	wincount, losecount, gamecount int
}

func (p *Player) NextHand() *hand {
	return p.Strategy.NextHand()
}

func (p *Player) Win() {
	p.wincount++
	p.gamecount++
}

func (p *Player) Lose() {
	p.losecount++
	p.gamecount++
}

func (p *Player) Even() {
	p.gamecount++
}
