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
	position Point
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
	return this.position
}

func (this Worm) busy() bool {
	return this.delay > 0
}

func (this Worm) move(w *World) {
	dir := rand.Intn(3)
	var newPos Point
	switch dir {
	case 0:
		newPos = this.position.Add(NORTH)
	case 1:
		newPos = this.position.Add(SOUTH)
	case 2:
		newPos = this.position.Add(EAST)
	case 3:
		newPos = this.position.Add(WEST)
	}
	if w.Contains(newPos) {
		this.position = newPos
		this.delay += 10
	} else {
		this.move(w)
	}
}
