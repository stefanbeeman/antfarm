package antfarm

const (
	SIMPLE  int = 1500
	COMPLEX int = 3000
)

type Action struct {
	delay    int
	complete func()
}

func (this Unit) generateWaitAction(duration int) Action {
	return Action{
		duration,
		func() {},
	}
}

func (this Unit) ready() bool {
	return this.currentAction.delay <= 1
}
