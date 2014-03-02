package af

type Actor interface {
	Location
	SetAction(Action)
}

type BasicActor struct {
	Location
	action Action
}

func (this BasicActor) SetAction(a Action) { this.action = a }

type Unit interface {
	Actor
	Thinker
	Mover
	memory() WorldState
	MovementCost(Location) (int,bool)
}

// func (this BasicMover) MovementCost(u, l Location) (int, bool) {
// 	cell := this.memory().GetCell(l)
// 	return 10, cell.getSolid()
// }