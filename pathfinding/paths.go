package pathfinding

import . "github.com/stefanbeeman/antfarm/common"

type PathStep struct {
  Point
  cost int
  best int
}

func (this PathStep) stepTo(pos Location, cost, h int) PathStep {
  newCost := this.cost + cost
  return PathStep{pos.AsPoint(), newCost, newCost + h}
}

type PathStack []Point
func (this PathStack) Empty() bool { return len(this) == 0 }
func (this *PathStack) Next() Point {
  orig := *this
  last := len(orig) - 1
  next := orig[last]
  *this = orig[:last]
  return next
}