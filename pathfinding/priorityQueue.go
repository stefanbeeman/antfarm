package pathfinding

import (
	"container/heap"
)

type PriorityQueue []PathStep

func (this PriorityQueue) Len() int           { return len(this) }
func (this PriorityQueue) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this PriorityQueue) Less(i, j int) bool {
	if this[i].best == this[j].best {
		return this[i].cost < this[j].cost
	} else {
		return this[i].best < this[j].best
	}
}

func (this *PriorityQueue) Push(x interface{}) {
	*this = append( *this, x.(PathStep) )
}

func (this *PriorityQueue) Pop() interface{} {
	old := *this
	n := len(old) - 1
	x := old[n]
	*this = old[0 : n]
	return x
}

func MakePriorityQueue() *PriorityQueue {
	q := &PriorityQueue{}
	heap.Init(q)
	return q
}