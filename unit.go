package af

type Actor interface {
	Location
	SetAction(Action)
	tic(WorldState)
}

type Unit interface {
	Actor
	Thinker
	Mover
	memory() WorldState
	MovementCost(Location) (int,bool)
}

type BasicUnit struct {
	Location
	Thinker
	Mover
	state WorldState
	action Action
}

func (this *BasicUnit) memory() WorldState { return this.state }
func (this *BasicUnit) SetAction(a Action) { this.action = a }

func (this *BasicUnit) tic(w WorldState) {
	if this.action.complete() {
		this.Think(this)
	}
	this.action.tic()
}

func (this *BasicUnit) MovementCost(l Location) (int, bool) {
	if this.memory().Contains(l) {
	  cell := this.memory().GetCell(l)
	  return 10, cell.getSolid()
	} else {
		return 0, true
	}
}


func MakeUnit(location, target Point, w WorldState) Unit {
	goal := &BasicGoal{target, 1}

	thinker := &BasicThinker{ []Goal{goal} }
	mover := &BasicMover{ MakeAStarAlg() }
	state := &OmniscientMemory{ w }
	result := &BasicUnit{location, thinker, mover, state, MakeWaitAction(0)}

	result.Init(result)
	return result
}