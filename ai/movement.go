package ai

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/pathfinding"
	"github.com/stefanbeeman/antfarm/storage"
)

type Mover interface {
	pathfinding.GoalDecider
	Move(Unit) Action
	initMover(Unit, storage.WorldState)
}

type BasicMover struct {
	pathfinding.GoalDecider
	alg pathfinding.MovementAlg
}

func (this *BasicMover) Move(u Unit) Action {
	next, valid := this.alg.NextStep(u, this.BestGoal())
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
		return u.memory().MoveCost(l)
	}

	this.alg.RegisterMoveCost(fn)
}

func MakeAStarMover() Mover {
	return &BasicMover{
		pathfinding.MakeGoalDecider(),
		pathfinding.MakeAStarAlg(),
	}
}