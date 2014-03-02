package antfarm

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/pathfinding"
	"github.com/stefanbeeman/antfarm/storage"
)

type Mover interface {
	pathfinding.MovementAlg
	Move(Unit) Action
	initMover(Unit, storage.WorldState)
}

func MakeAStarMover() BasicMover {
	return BasicMover{
		pathfinding.MakeAStarAlg(),
	}
}

type BasicMover struct {
	pathfinding.MovementAlg
}

func (this *BasicMover) MoveCost(l Location) (int, bool) {
	// cell := this.memory().GetCell(l)
	return 10, false
}

func (this *BasicMover) Move(u Unit) Action {
	next, valid := this.NextStep(u)
	if !valid {
		return MakeWaitAction(0)
	}
	return MakeMoveAction(u, next, 10)
}

func (this *BasicMover) initMover(u Unit, w storage.WorldState) {
	fn := func(l Location) (int, bool) {
		if !w.Contains(l) {
			return 0, true
		}
		return this.MoveCost(l)
	}

	this.RegisterMoveCost(fn)
}
