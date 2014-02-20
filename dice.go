package af

import "math/rand"

type dice struct{}

func (this dice) d6() int {
	return (rand.Intn(6) + 1)
}

func (this dice) count(rolls []int, fn func(int) bool) int {
	total := 0
	for roll := range rolls {
		if fn(roll) {
			total++
		}
	}
	return total
}

func (this dice) countHits(rolls []int) int {
	return this.count(rolls, func(roll int) bool {
		return roll > 4
	})
}

func (this dice) countOnes(rolls []int) int {
	return this.count(rolls, func(roll int) bool {
		return roll <= 1
	})
}

func (this dice) isGlitch(rolls []int) bool {
	dice := len(rolls)
	ones := this.countOnes(rolls)
	return ones >= (dice / 2)
}

func (this dice) applyLimit(hits int, limit int) int {
	if hits > limit {
		return limit
	} else {
		return hits
	}
}

func (this dice) rollDice(dice int, limit int) (int, bool) {
	rolls := make([]int, dice)
	for i := range rolls {
		rolls[i] = this.d6()
	}
	hits := this.applyLimit(this.countHits(rolls), limit)
	glitch := this.isGlitch(rolls)
	return hits, glitch
}

// func (this dice) successTest(dice int, limit int, threshold int) (bool, bool) {
//  hits, glitch := this.rollDice(dice, limit)
//  return (hits >= threshold), glitch
// }

// func (this dice) opposedTest(myDice int, myLimit int, yourDice int, yourLimit int) (int, bool, bool) {
//  myHits, myGlitch := this.rollDice(myDice, myLimit)
//  yourHits, yourGlitch := this.rollDice(yourDice, yourLimit)
//  return (myHits - yourHits), myGlitch, yourGlitch
// }

var afd = dice{}
