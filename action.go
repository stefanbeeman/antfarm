package antfarm

import "fmt"

const (
	SIMPLE  int = 1500
	COMPLEX int = 3000
)

type Action interface {
	complete() bool
	tic()
}

type BasicAction struct {
	currentDelay int
	fnComplete   func()
}

func (this *BasicAction) tic() {
	this.currentDelay--
	if this.currentDelay < 1 {
		this.fnComplete()
	}
}

func (this BasicAction) complete() bool {
	return this.currentDelay < 1
}

func makeWaitAction(duration int) Action {
	act := BasicAction{
		duration,
		func() {},
	}
	return &act
}
