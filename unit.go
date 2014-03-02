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

type BasicUnit struct {
	Actor
	Thinker
	Mover
	state WorldState
}

func (this BasicUnit) memory() WorldState { return this.state }

func (this BasicUnit) MovementCost(l Location) (int, bool) {
	cell := this.memory().GetCell(l)
	return 10, cell.getSolid()
}


func MakeUnit(location, target Point, w WorldState) Unit {
	goal := &BasicGoal{target, 1}

	actor := BasicActor{ location, MakeWaitAction(0) }
	thinker := BasicThinker{ []Goal{goal} }
	mover := BasicMover{ MakeAStarAlg() }
	state := OmniscientMemory{ w }
	result := BasicUnit{actor, thinker, mover, state}

	result.Init(result)
	return result
}