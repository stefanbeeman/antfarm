package pathfinding

import (
  . "github.com/stefanbeeman/antfarm/common"
  "container/heap"
)

type AStarAlg struct {
  path PathStack
  MoveCost func(Location) (int, bool) // movement cost, blocked
}

func (this *AStarAlg) NextStep(start, goal Location) (Location, bool) {
  if this.path.Empty() {
    this.findPath(start, goal)
  }

  if !this.path.Empty() {
    next := this.path.Next()
    if _, blocked := this.MoveCost(next); !blocked  {
      return next, true
    }
  }

  return Point{}, false
}

func (this *AStarAlg) findPath(start, goal Location) bool {
  startPoint := start.AsPoint()
  q := MakeAStarQueue()
  q.Close(startPoint)

  current := PathStep{startPoint, 0, this.h(start, goal)}
  for !current.At(goal) {
    for _, next := range current.Neighbors() {
      if cost, blocked := this.MoveCost(next); !blocked {
        nextStep := current.stepTo(next, cost, this.h(next, goal))
        q.Insert( current, nextStep )
      }
    }

    if (q.Len() == 0) {
      return false
    }
    current = q.Next()
  }

  this.path = q.Rewind(current.AsPoint(), startPoint)
  return !this.path.Empty()
}

func (this *AStarAlg) h(start, goal Location) int { return 10 * start.DistanceTo(goal) }
func (this *AStarAlg) RegisterMoveCost(fn func(Location) (int, bool)) { this.MoveCost = fn }


type AStarQueue struct {
  q *PriorityQueue
  closedSet map[Point]Point
}

func (this *AStarQueue) Len() int { return this.q.Len() }
func (this *AStarQueue) Next() PathStep { return heap.Pop(this.q).(PathStep) }
func (this *AStarQueue) Close(point Point) { this.closedSet[point] = point }

func (this *AStarQueue) Insert(from, to PathStep) bool {
  fromPos, toPos := from.AsPoint(), to.AsPoint()
  if _, seen := this.closedSet[toPos]; !seen {
    heap.Push(this.q, to)
    this.closedSet[toPos] = fromPos
    return true
  }
  return false
}

func (this *AStarQueue) Rewind(end, start Point) PathStack {
  result := PathStack{}
  for next := end; !next.At(start); next = this.closedSet[next] {
    result = append(result, next)
  }
  return result
}

func MakeAStarQueue() *AStarQueue {
  closedSet := make(map[Point]Point)
  result := &AStarQueue{MakePriorityQueue(), closedSet}
  return result
}

func MakeAStarAlg() MovementAlg {
  return &AStarAlg{
    []Point{},
    nil,
  }
}