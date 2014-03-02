package antfarm

import (
	"github.com/stefanbeeman/antfarm/pathfinding"
)

type Thinker interface {
	Init(Unit)
	Think(Unit)
}

type BasicThinker struct {
	goals []pathfinding.Goal
}

func (this *BasicThinker) Init(u Unit) {
	u.AddGoals(this.goals)
}

func (this *BasicThinker) Think(u Unit) {
	u.SetAction(u.Move(u))
}
