package af

import "math/rand"

type DiceRoller interface {
	d6() int
	d10() int
	rollDice(int, int) (int, bool)
}

type BasicDiceRoller struct{}

func (this BasicDiceRoller) d6() int {
	return (rand.Intn(6) + 1)
}

func (this BasicDiceRoller) d10() int {
	die := (rand.Intn(10) + 1)
	if die == 10 {
		return die + this.d10()
	} else {
		return die
	}
}

func (this BasicDiceRoller) count(rolls []int, fn func(int) bool) int {
	total := 0
	for roll := range rolls {
		if fn(roll) {
			total++
		}
	}
	return total
}

func (this BasicDiceRoller) countHits(rolls []int, tn int) int {
	return this.count(rolls, func(roll int) bool {
		return roll >= tn
	})
}

func (this BasicDiceRoller) countGlitches(rolls []int) int {
	return this.count(rolls, func(roll int) bool {
		return roll <= 2
	})
}

func (this BasicDiceRoller) isGlitch(rolls []int) bool {
	dice := len(rolls)
	glitches := this.countGlitches(rolls)
	return glitches >= (dice / 2)
}

func (this BasicDiceRoller) rollDice(BasicDiceRoller int, tn int) (int, bool) {
	rolls := make([]int, BasicDiceRoller)
	for i := range rolls {
		rolls[i] = this.d10()
	}
	hits := this.countHits(rolls, tn)
	glitch := this.isGlitch(rolls)
	return hits, glitch
}

var dice = BasicDiceRoller{}
