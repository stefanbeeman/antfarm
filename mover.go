package antfarm

import (
	"math/rand"
)

type Mover interface {
	where() Point
	move(*World)
	moveDelay(Cell) int
}

type RandomMover struct {
	Location Point
	Speed    int
}

func (this RandomMover) where() Point {
	return this.Location
}

func (this RandomMover) moveDelay(through Cell) int {
	return this.Speed
}

func (this *RandomMover) move(w *World) {
	dir := rand.Intn(3)
	var newPos Point
	switch dir {
	case 0:
		newPos = this.Location.Add(NORTH)
	case 1:
		newPos = this.Location.Add(SOUTH)
	case 2:
		newPos = this.Location.Add(EAST)
	case 3:
		newPos = this.Location.Add(WEST)
	}
	if w.Contains(newPos) {
		this.Location = newPos
	} else {
		this.move(w)
	}
}
