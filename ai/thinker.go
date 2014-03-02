package ai

import (
	"github.com/stefanbeeman/antfarm/pathfinding"
)

type Thinker interface {
	Think(Unit)
  initThinker(Unit)
}

type BasicThinker struct {
	goals []pathfinding.Goal
}

func (this *BasicThinker) initThinker(u Unit) {
	u.AddGoals(this.goals)
}

func (this *BasicThinker) Think(u Unit) {
	u.SetAction(u.Move(u))
}
