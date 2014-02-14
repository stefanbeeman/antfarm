package antfarm

import (
	"fmt"
	"math/rand"
)

type Unit struct {
	Actor
	Mover
	currentTask Task
	desires     DesireQueue
	world       *World
	speed       int
}

func (this *Unit) tic() {
	this.Actor.tic()
	if this.desires.decide() || this.currentTask.invalid() {
		this.currentTask = this.generateTask(this.desires.next())
	}
	if this.ready() {
		this.currentAction = this.currentTask.next()
	}
}

func (this *Unit) generateTask(d Desire) Task {
	switch d.verb {
	case "wander":
		target := this.world.Random()
		return this.generateMoveToTask(target)
	default:
		return this.generateComaTask()
	}
}

func (this *Unit) generateMoveToTask(target Point) Task {
	complete := func() bool {
		if this.Position == target {
			return true
		} else {
			return false
		}
	}
	invalid := func() bool {
		// At some point this should check to see if it's possible for the unit to move to that point at all.
		return false
	}
	next := func() Action {
		dir := this.Position.vectorTo(target)
		if dir.X != 0 && dir.Y != 0 {
			//We need to choose the faster direction, here. For now, we pick at random.
			coin := rand.Intn(2)
			if coin == 0 {
				dir.X = 0
			} else {
				dir.Y = 0
			}
		}
		moveComplete := func() {
			this.Move(dir)
		}
		return Action{this.speed, moveComplete}
	}
	return Task{complete, invalid, next}
}

func (this *Unit) generateComaTask() Task {
	return Task{
		func() bool {
			return false
		},
		func() bool {
			return false
		},
		func() Action {
			return Action{this.speed, func() {
				fmt.Println("zzzzzzzz")
			}}
		},
	}
}
