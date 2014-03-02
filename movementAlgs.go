package af

type MovementAlg interface {
  GoalDecider
  Move(Unit) Action
}

type GoalDecider interface {
  AddGoals([]Goal)
  RemoveGoals([]Goal)
  BestGoal() Goal
  H(Location) int
}

type BasicGoalDecider struct {
  goals []Goal
}

func (this *BasicGoalDecider) AddGoals(goals []Goal) {
  this.goals = append( this.goals, goals... )
}

func (this *BasicGoalDecider) RemoveGoals(goals []Goal) {
  removed := make(map[Goal]bool)
  for _, g := range goals {
    removed[g] = true
  }
  modified := []Goal{}
  for _, g := range this.goals {
    if _, ok := removed[g]; ok {
      modified = append( modified, g )
    }
  }
  this.goals = modified
}

func (this *BasicGoalDecider) BestGoal() Goal  {
  return this.goals[0]
}

func (this *BasicGoalDecider) H(p Location) int {
  return p.DistanceTo( this.BestGoal() )
}

func MakeGoalDecider() GoalDecider { return &BasicGoalDecider{[]Goal{}} }



type PathStep struct {
  Point
  cost int
  best int
}

func (this PathStep) stepTo(pos Location, cost, h int) PathStep {
  newCost := this.cost + cost
  return PathStep{pos.AsPoint(), newCost, newCost + h}
}



type AStarAlg struct {
  GoalDecider
  path []Point
}

func (this *AStarAlg) Move(u Unit) Action {
  next, valid := this.nextPlannedStep(u)
  if !valid {
    if !this.plan(u) {
      return MakeWaitAction(0)
    }
    next,_ = this.nextPlannedStep(u)
  }
  return &BasicAction{0, 10, func(){u.SetPosition(next)}}
}

func (this *AStarAlg) nextPlannedStep(u Unit) (Location, bool) {
  l := len(this.path)
  if l > 0 {
    next := this.path[l-1]
    this.path = this.path[:l-1]
    if _, valid := u.MovementCost(next); valid  {
      return next, true
    }
  }
  return Point{}, false
}

func (this *AStarAlg) plan(u Unit) bool {
  start := u.AsPoint()
  goal := u.BestGoal()

  q := MakeAStarQueue()
  q.Close(start)

  current := PathStep{start, 0, this.H(start)}
  for !current.At(goal) {
    for _, next := range current.Neighbors() {
      if cost, traversable := u.MovementCost(next); traversable {
        q.Insert( current, current.stepTo(next, cost, this.H(next)) )
      }
    }

    if (q.Len() == 0) {
      return false
    }
    current = q.Next()
  }

  this.path = q.Rewind(current.AsPoint(), start)
  return true
}

func MakeAStarAlg() MovementAlg { return &AStarAlg{MakeGoalDecider(), []Point{}} }