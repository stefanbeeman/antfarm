package antfarm

func (this BasicGoal) Weight() int { return this.weight }

type Thinker interface {
  Init(Unit)
  Think(Unit)
}

type BasicThinker struct{
  goals []Goal
}

func (this *BasicThinker) Init(u Unit) {
  u.AddGoals( this.goals )
}

func (this *BasicThinker) Think(u Unit) {
  u.SetAction( u.Move(u) )
}