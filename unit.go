package antfarm

type Unit struct {
	Actor
	Mover
	currentTask Task
	desires     DesireQueue
}

func (this *Unit) tic(w *World) {
	this.Actor.tic()
	if this.desires.decide() || this.currentTask.invalid() {
		this.currentTask = this.generateTask(this.desires.next(), w)
	}
	if this.ready() {
		this.currentAction = this.currentTask.next(w)
	}
}

func (this *Unit) generateTask(d Desire) Task {
	switch d.verb {
	case "kill":
		return "foo"
	case "move":
		return "bar"
	}
}
