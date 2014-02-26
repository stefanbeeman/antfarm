package af

import "fmt"

type Actor interface {
	tic(World)
}

type Unit interface {
	Actor
}

type BasicUnit struct {
	Name          string
	Species       string
	currentAction Action
	currentTask   Task
	Position      Point
	thinker       *BasicThinker
	mover         Mover
	memory        World
}

func (this *BasicUnit) tic(w World) {
	this.currentAction.tic()

	if this.currentAction.complete() {
		if this.currentTask.complete() {
			this.currentTask = this.thinker.think(this, w)
		}
		fmt.Println("I'm at ", this.Position)
		this.currentAction = this.currentTask.getAction()
	}
}

func (this BasicUnit) moveCost(p Point) (float64, bool) {
	solid := this.memory.get(p).getSolid()
	return 1, !solid
}

func makeWorm(where Point, w World) Actor {
	m := new(AStarWalker)
	worm := BasicUnit{
		"Wormy the Worm",
		"Worm",
		makeWaitAction(1),
		makeDummyTask(),
		where,
		new(BasicThinker),
		m,
		w,
	}
	return &worm
}
