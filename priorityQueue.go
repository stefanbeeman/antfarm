package af

import (
	"container/heap"
)

type AStarQueue struct {
	q []PathStep
	closedSet map[Point]Point
}

func MakeAStarQueue() *AStarQueue {
	closedSet := make(map[Point]Point)
	q := []PathStep{}
	result := &AStarQueue{q, closedSet}
	heap.Init(result)
	return result
}

func (this *AStarQueue) Insert(from, to PathStep) bool {
	fromPos, toPos := from.AsPoint(), to.AsPoint()
	if _, seen := this.closedSet[toPos]; !seen {
		heap.Push(this, to)
		this.closedSet[toPos] = fromPos
		return true
	}
	return false
}

func (this *AStarQueue) Next() PathStep {
	return heap.Pop(this).(PathStep)
}

func (this *AStarQueue) Close(point Point) {
	this.closedSet[point] = point
}

func (this *AStarQueue) Rewind(end, start Point) []Point {
	result := []Point{}
	for next := end; !next.At(start); next = this.closedSet[next] {
		result = append(result, next)
	}
	return result
}

func (this *AStarQueue) Len() int           { return len(this.q) }
func (this *AStarQueue) Swap(i, j int)      { this.q[i], this.q[j] = this.q[j], this.q[i] }

func (this *AStarQueue) Less(i, j int) bool {
	if this.q[i].best == this.q[j].best {
		return this.q[i].cost < this.q[j].cost
	} else {
		return this.q[i].best < this.q[j].best
	}
}

func (this *AStarQueue) Push(x interface{}) {
	this.q = append( this.q, x.(PathStep) )
}

func (this AStarQueue) Pop() interface{} {
	old := this.q
	n := len(old)
	x := old[n-1]
	this.q = old[0 : n-1]
	return x
}
