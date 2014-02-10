package antfarm

import (
	"math/rand"
)

type Unit interface {
	Act(*World) Unit
	Where() Point
	busy() bool
	move(*World)
}

type Worm struct {
	Position Point
	delay    int
}

func (this Worm) Act(w *World) Unit {
	if !this.busy() {
		this.move(w)
	} else {
		this.delay += -1
	}
	return this
}

func (this Worm) Where() Point {
	return this.Position
}

func (this Worm) busy() bool {
	return this.delay > 0
}

func (this Worm) move(w *World) {
	dir := rand.Intn(3)
	var newPos Point
	switch dir {
	case 0:
		newPos = this.Position.Add(NORTH)
	case 1:
		newPos = this.Position.Add(SOUTH)
	case 2:
		newPos = this.Position.Add(EAST)
	case 3:
		newPos = this.Position.Add(WEST)
	}
	if w.Contains(newPos) {
		this.Position = newPos
		this.delay += 10
	} else {
		this.move(w)
	}
}
