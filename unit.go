package antfarm

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/pathfinding"
	"github.com/stefanbeeman/antfarm/storage"
)

type Actor interface {
	MutableLocation
	SetAction(Action)
	tic(storage.WorldState)
}

type Unit interface {
	Actor
	Thinker
	Mover
	memory() storage.WorldState
}

type BasicUnit struct {
	MutableLocation
	Thinker
	Mover
	state  storage.WorldState
	action Action
}

func (this *BasicUnit) memory() storage.WorldState { return this.state }
func (this *BasicUnit) SetAction(a Action)         { this.action = a }

func (this *BasicUnit) tic(w storage.WorldState) {
	if this.action.complete() {
		this.Think(this)
	}
	this.action.tic()
}

func MakeUnit(location, target Point, w storage.WorldState) Unit {
	goal := &pathfinding.BasicGoal{target, 1}
	thinker := &BasicThinker{[]pathfinding.Goal{goal}}
	mover := &BasicMover{pathfinding.MakeAStarAlg()}
	state := &OmniscientMemory{w}
	result := &BasicUnit{&location, thinker, mover, state, MakeWaitAction(0)}

	result.Init(result)
	return result
}
