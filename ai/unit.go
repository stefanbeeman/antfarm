package ai

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/pathfinding"
	"github.com/stefanbeeman/antfarm/storage"
)

type Actor interface {
	MutableLocation
	SetAction(Action)
	Tic(storage.WorldState)
	Display() Display
}

type Unit interface {
	Actor
	Thinker
	Mover
	memory() Memory
	Init(storage.WorldState)
}

type BasicUnit struct {
	MutableLocation
	Thinker
	Mover
	state  Memory
	action Action
}

func (this *BasicUnit) memory() Memory     { return this.state }
func (this *BasicUnit) SetAction(a Action) { this.action = a }
func (this *BasicUnit) Tic(w storage.WorldState) {
	if this.action.complete() {
		this.Think(this)
	}
	this.action.tic()
}

func (this *BasicUnit) Init(world storage.WorldState) {
	this.initThinker(this)
	this.initMover(this, world)
}

type DisplayActor struct {
	X    int
	Y    int
	Tile string
}

func (this BasicUnit) Display() Display {
	tile := "Worm"
	x, y := this.Coords()
	return DisplayActor{x, y, tile}
}

func MakeUnit(location, target Point, w storage.WorldState) Unit {
	goal := &pathfinding.BasicGoal{target, 1}
	thinker := &BasicThinker{[]pathfinding.Goal{goal}}
	mover := MakeAStarMover()
	state := &OmniscientMemory{w}
	result := &BasicUnit{&location, thinker, mover, state, MakeWaitAction(0)}

	result.Init(w)
	return result
}
