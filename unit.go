package af

import "fmt"

type Actor interface {
	func SetAction(Action)
}

type BasicActor struct {
	Location
	action Action
}

func (this BasicActor) SetAction(a Action) { this.Action = a }

type Unit interface {
	Actor
	Thinker
	Mover
	memory() WorldState
}
