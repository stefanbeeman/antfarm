package antfarm

type Actor interface {
	tic(*World)
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
	mover         *RandomWalker
}

func (this BasicUnit) tic(w *World) {
	this.currentAction.tic()
	if this.currentAction.complete() {
		this.currentTask = this.thinker.think(&this, w)
		this.currentAction = this.currentTask.getAction()
	}
}

func makeWorm(where Point) Actor {
	return BasicUnit{
		"Wormy the Worm",
		"Worm",
		makeWaitAction(1),
		makeDummyTask(),
		where,
		new(BasicThinker),
		new(RandomWalker),
	}
}
