package antfarm

type Unit interface {
	Actor
	ready() bool
}

type BasicUnit struct {
	Name          string
	Species       string
	currentAction Action
	Position      Point
	thinker       Thinker
	mover         Mover
}

func (this BasicUnit) tic() {}

func (this BasicUnit) ready() bool {
	return this.currentAction.delay <= 1
}
