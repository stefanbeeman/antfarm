package antfarm

import (
  "github.com/stefanbeeman/antfarm/pathfinding"
)

type Mover interface {
  Move(Unit) Action
  InitMover(Unit, WorldState)
}

func MakeAStarMover() BasicMover {
  return BasicMover{
    pathfinding.MakeAStarAlg(),
  }
}

type BasicMover struct {
	alg MovementAlg
}

func (this *BasicMover) MoveCost(l Location) (int, bool) {
  cell := this.memory().GetCell(l)
  return 10, cell.getSolid()
}

func (this *BasicMover) Move(u Unit) Action {
  next, valid := this.alg.NextPlannedStep(u)
  if !valid { return MakeWaitAction(0) }
  return MakeMoveAction(u, next, 10)
}

func (this *BasicMover) InitMover(u Unit, w World) {
  fn := func(l Location) (int, bool) {
    if !w.Contains(l) { return 0, true }
    return this.MoveCost(l)
  }

  this.alg.RegisterMoveCost( fn )
}