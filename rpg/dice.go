package rpg

import (
	"math/rand"
)

type Roll interface {
	GetDice() int
	GetTN() int
	GetShade() int
}

type BasicRoll struct {
	Dice  int
	TN    int
	Shade int
}

func (this BasicRoll) GetDice() int {
	return this.Dice
}

func (this BasicRoll) GetTN() int {
	return this.TN
}

func (this BasicRoll) GetShade() int {
	return this.Shade
}

type DiceRoller interface {
	DN(int) int
	NDN(int, int) int
	D6() int
	D10() int
	RollDice(Roll) int
}

type BasicDiceRoller struct{}

func (this BasicDiceRoller) DN(sides int) int {
	return (rand.Intn(sides) + 1)
}

func (this BasicDiceRoller) NDN(n int, sides int) int {
	total := 0
	for n > 0 {
		total += this.DN(sides)
		n--
	}
	return total
}

func (this BasicDiceRoller) D6() {
	return this.DN(6)
}

func (this BasicDiceRoller) D10() {
	return this.DN(10)
}

func (this BasicDiceRoller) count(rolls []int, fn func(int) bool) int {
	total := 0
	for _, roll := range rolls {
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
		return roll < 3
	})
}

func (this BasicDiceRoller) shadeGrey(rolls []int) []int {
	more := this.count(rolls, func(roll int) bool {
		return roll > 10
	})
	extra := make([]int, 0)
	for more > 0 {
		die := this.D12()
		extra = append(extra, die)
		more--
	}
	return append(rolls, extra...)
}

func (this BasicDiceRoller) shadeWhite(rolls []int) []int {
	more := this.count(rolls, func(roll int) bool {
		return roll > 10
	})
	extra := make([]int, 0)
	for more > 0 {
		die := this.D12()
		extra = append(extra, die)
		more--
	}
	if len(extra) > 0 {
		extra = this.shadeWhite(extra)
	}
	return append(rolls, extra...)
}

func (this BasicDiceRoller) RollDice(what Roll) int {
	dice := what.GetDice()
	tn := what.GetTN()
	shade := what.GetShade()
	rolls := make([]int, dice)
	for i := range rolls {
		rolls[i] = this.D12()
	}
	if shade <= WHITE {
		rolls = this.shadeWhite(rolls)
	} else if shade == GRAY {
		rolls = this.shadeGrey(rolls)
	}
	hits := this.countHits(rolls, tn)
	glitches := this.countGlitches(rolls)
	return hits, glitches
}

var Dice = BasicDiceRoller{}
