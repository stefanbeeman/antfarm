package af

import (
	"container/heap"
)

type AStarQueue struct {
	q []PathStep
	closedSet map[Location]Location
}

func MakeAStarQueue() AStarQueue {
	closedSet := make(map[Point]Point)
	q := &PathStepPriorityQueue{}
	heap.Init(q)
	return AStarQueue{q, closedSet}
}

func (this AStarQueue) Insert(from, to PathStep) bool {
	fromPos, toPos := from.Position(), to.Position()
	if _, seen := this.closedSet[toPos]; !seen {
		heap.Push(this.q, to)
		this.closedSet[toPos] = fromPos
		return true
	}
	return false
}

func (this AStarQueue) Next() PathStep {
	return heap.Pop(this.q).(PathStep)
}

func (this AStarQueue) Close(point Point) {
	closedSet[point] = point
}

func (this AStarQueue) Rewind(end, start Point) []Point {
	result := []Point{}
	step := this.closedSet(end.AsPoint())
	for next := end; !next.At(start); next = closedSet[next] {
		result = append(result, next)
	}
	return result
}

func (this AStarQueue) Len() int           { return len(this.q) }
func (this AStarQueue) Swap(i, j int)      { this.q[i], this.q[j] = this.q[j], this.q[i] }

func (this AStarQueue) Less(i, j int) bool {
	if this.q[i].best == this.q[j].best {
		return this.q[i].cost < this.q[j].cost
	} else {
		return this.q[i].best < this.q[j].best
	}
}

func (this AStarQueue) Push(x interface{}) {
	this.q = append( this.q, x.(MoveStep) )
}

func (this AStarQueue) Pop() interface{} {
	old := this.q
	n := len(old)
	x := old[n-1]
	this.q = old[0 : n-1]
	return x
}
