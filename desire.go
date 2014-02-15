package antfarm

import "sort"

type Verb string

type Desire struct {
	verb Verb
	noun interface{}
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

func (this *DesireQueue) add(d Desire) {
	this.queue = append(this.queue, d)
}

func (this *DesireQueue) decide() bool {
	top := this.queue[0]
	sort.Sort(this)
	return this.queue[0] == top
}

func (this DesireQueue) next() Desire {
	return this.queue[0]
}
