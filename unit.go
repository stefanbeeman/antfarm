package antfarm

type Unit struct {
	world         *World
	Name          string
	Species       string
	Position      Point
	LandSpeed     Speed
	currentAction Action
	currentTask   Task
	desires       DesireQueue
	Statline
}

func (this *Unit) tic(w *World) {
	this.currentAction.delay--
	if this.ready {
		this.currentAction.complete()
		this.think(w)
	}
}

func (this *Unit) think(w *World) {
	newDesire := false
	top := this.desires.next()
	bottom := this.desires.decide()
	if top != bottom {
		newDesire = true
	}
	if newDesire || this.currentTask.complete() {
		this.currentTask = this.generateTask()
	}
	this.currentAction = this.currentTask.next()
}
