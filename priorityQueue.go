package af

type PositionHeap struct {
	queue []positionHeapEntry
}

type positionHeapEntry struct {
	p      Point
	weight float64
}

func MakePositionHeap() *PositionHeap {
	result := new(PositionHeap)
	result.queue = []positionHeapEntry{}
	return result
}

func (this *PositionHeap) Size() int {
	return len(this.queue)
}

func (this *PositionHeap) Push(p Point, weight float64) {
	this.queue = append(this.queue, positionHeapEntry{p, weight})
	this.bubbleUp()
}

func (this *PositionHeap) Pop() (Point, float64) {
	result := this.queue[0]
	n := len(this.queue) - 1
	this.queue[0], this.queue[n] = this.queue[n], this.queue[0]
	this.queue = this.queue[0:n]
	this.sinkDown()
	return result.p, result.weight
}

func (this *PositionHeap) resize() {
	l := len(this.queue)
	newQueue := make([]positionHeapEntry, l, 2*l+1)
	copy(newQueue, this.queue)
	this.queue = newQueue
}

func (this *PositionHeap) bubbleUp() {
	slice := this.queue
	n := len(slice) - 1
	p := parent(n)
	for (p >= 0) && (slice[n].weight < slice[p].weight) {
		slice[n], slice[p] = slice[p], slice[n]
		n = p
	}
}

func (this *PositionHeap) sinkDown() {
	slice := this.queue
	n := 0
	c := this.smallerChild(n)

	for n != c && slice[n].weight > slice[c].weight {
		slice[n], slice[c] = slice[c], slice[n]
		n = c
		c = this.smallerChild(n)
	}
}

func (this *PositionHeap) smallerChild(n int) int {
	left := child(n)
	right := left + 1
	capac := len(this.queue)

	if right == capac {
		return left
	} else if right < capac {
		if this.queue[left].weight < this.queue[right].weight {
			return left
		} else {
			return right
		}
	} else {
		return n
	}
}

func parent(n int) int {
	return (n - 1) / 2
}

func child(n int) int {
	return 2*n + 1
}
