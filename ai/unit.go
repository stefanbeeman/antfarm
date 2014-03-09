package ai

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/pathfinding"
	"github.com/stefanbeeman/antfarm/world"
)

type Actor interface {
	MutableLocation
	SetAction(Action)
	Tic(world.WorldState)
	Display() Display
}

type Unit interface {
	Actor
	Thinker
	Mover
	memory() Memory
	Init(world.WorldState)
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
func (this *BasicUnit) Tic(w world.WorldState) {
	if this.action.complete() {
		this.Think(this)
	}
	this.action.tic()
}

func (this *BasicUnit) Init(world world.WorldState) {
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

func MakeUnit(location, target Point, w world.WorldState) Unit {
	goal := &pathfinding.BasicGoal{target, 1}
	thinker := &BasicThinker{[]pathfinding.Goal{goal}}
	mover := MakeAStarMover()
	state := &OmniscientMemory{w}
	result := &BasicUnit{&location, thinker, mover, state, MakeWaitAction(0)}

	result.Init(w)
	return result
}
