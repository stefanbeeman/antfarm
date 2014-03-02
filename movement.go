package af

import (
	"container/heap"
)

type Mover interface {
	MovementAlg
	MovementCost(Unit, Location) (int, bool)
}

type BasicMover struct {
	MovementAlg
	world *World
}

function (this BasicMover) MovementCost(l Location) (int, bool) {
	cell := this.memory().GetCell(l)
	return 10, cell.getSolid()
}
