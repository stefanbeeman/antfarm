package af

import "math/rand"

func d6() int {
	return (rand.Intn(6) + 1)
}

func count(rolls []int, fn func(int) bool) int {
	total := 0
	for roll := range rolls {
		if fn(roll) {
			total++
		}
	}
	return total
}

func countHits(rolls []int) int {
	return count(rolls, func(roll int) bool {
		return roll > 4
	})
}

func countOnes(rolls []int) int {
	return count(rolls, func(roll int) bool {
		return roll <= 1
	})
}

func isGlitch(rolls []int) bool {
	dice := len(rolls)
	ones := countOnes(rolls)
	return ones >= (dice / 2)
}

func applyLimit(hits int, limit int) int {
	if hits > limit {
		return limit
	} else {
		return hits
	}
}

func rollDice(dice int, limit int) (int, bool) {
	rolls := make([]int, dice)
	for i := range rolls {
		rolls[i] = d6()
	}
	hits := applyLimit(countHits(rolls), limit)
	glitch := isGlitch(rolls)
	return hits, glitch
}

func successTest(dice int, limit int, threshold int) (bool, bool) {
	hits, glitch := rollDice(dice, limit)
	return (hits >= threshold), glitch
}

func opposedTest(myDice int, myLimit int, yourDice int, yourLimit int) (int, bool, bool) {
	myHits, myGlitch := rollDice(myDice, myLimit)
	yourHits, yourGlitch := rollDice(yourDice, yourLimit)
	return (myHits - yourHits), myGlitch, yourGlitch
}
