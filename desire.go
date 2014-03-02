package antfarm

import . "github.com/stefanbeeman/antfarm/common"

import "sort"

type Verb string

type Desire struct {
	verb Verb
	noun interface{}
}

type DesirePoint struct {
	Position Point
	Weight float64
}

type DesireQueue struct {
	queue []Desire
	weigh func(Desire) int
}

func (this DesireQueue) Len() int {
	return len(this.queue)
}

func (this DesireQueue) Less(i int, j int) bool {
	return this.weigh(this.queue[i]) > this.weigh(this.queue[j])
}

func (this *DesireQueue) Swap(i int, j int) {
	this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *DesireQueue) add(d Desire) Desire {
	this.queue = append(this.queue, d)
	return this.decide()
}

func (this *DesireQueue) decide() Desire {
	sort.Sort(this)
	return this.queue[0]
}

func (this DesireQueue) next() Desire {
	return this.queue[0]
}

func makeDesireQueue() DesireQueue {
	return DesireQueue{
		make([]Desire, 0),
		func(d Desire) int { return 0 },
	}
}
