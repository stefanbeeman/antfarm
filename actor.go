package antfarm

type Action struct {
	delay    int
	complete func()
}

type Actor struct {
	currentAction Action
}

func (this *Actor) act(a Action) {
	this.currentAction = a
}

func (this *Actor) react(a Action) {
	this.currentAction.delay = this.currentAction.delay + a.delay
	a.complete()
}

func (this *Actor) ready() bool {
	return this.currentAction.delay < 1
}

func (this *Actor) tic() {
	if this.ready() {
		this.currentAction.complete()
	} else {
		this.currentAction.delay--
	}
}
