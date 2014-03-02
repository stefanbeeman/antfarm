
type MovementAlg interface {
  GoalDecider
  Move(Unit) Action
}

type GoalDecider interface {
  AddGoals([]Goal)
  RemoveGoals([]Goal)
  BestGoal(Unit) Goal
  H(Location) int
}

type BasicGoalDecider struct {
  goals []Goal
}

func (this BasicGoalDecider) AddGoals(goals Goal[]) {
  this.goals = append( this.goals, goals... )
}

func (this BasicGoalDecider) RemoveGoals(goals Goal[]) {
  removed := make(map[Goal]bool)
  for _, g := range goals {
    removed[g] = true
  }
  modified := []Goal{}
  for _, g := range this.goals {
    if _, ok := removed[g], ok {
      modified = append( modified, g )
    }
  }
  this.goals = modified
}

func (this BasicGoalDecider) BestGoal(u Unit) Goal  {
  return this.goals[0]
}

func (this BasicGoalDecider) H(p Location) int {
  return p.DistanceTo( this.BestGoal() )
}



type PathStep struct {
  Point
  cost int
  best int
}

func (this PathStep) to(pos Point, cost, h) PathStep {
  newCost := this.cost + cost
  return PathStep{pos, newCost, newCost + h}
}



type AStarAlg struct {
  GoalDecider
  path[] Point
}

func (this AStarAlg) Move(u Unit) Action {
  next, valid := this.next()
  if !valid {
    if !plan(u) {
      return Action{"Bad things"}
    }
    next,_ = this.next()
  }
  return Action{next}
}

func (this AStarAlg) next(u Unit) (Location, bool) {
  l := len(this.path)
  if l > 0 {
    next := this.path[l-1]
    this.path = this.path[:i]
    if _, valid := u.MovementCost(next); valid  {
      return next, true
    }
  }
  return Point{}, false
}

func (this AStarAlg) plan(u Unit) bool {
  start := this.AsPoint()
  goal := BestGoal()

  q := MakeAStarQueue()
  q.Close(start)

  for current = PathStep{start, 0, this.H(start)}; !current.At(goal); current = q.Next() {
    for _, next := range current.Neighbors() {
      if cost, traversable := u.MovementCost(p); traversable {
        q.Insert( current, current.to(next) )
      }
    }

    if (q.Len() == 0) {
      return false
    }
  }

  this.path = q.Rewind(current.AsPoint(), start)
  return true
}