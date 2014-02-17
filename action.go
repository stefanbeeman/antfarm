package antfarm

const (
	SIMPLE  int = 1500
	COMPLEX int = 3000
)

type Actor interface {
	tic(*World)
}

type Action interface {
	complete()
}

type BasicAction struct {
	delay      int
	fnComplete func()
}

func (this BasicAction) complete() {
	this.fnComplete()
}

func makeWaitAction(duration int) Action {
	return BasicAction{
		duration,
		func() {},
	}
}
