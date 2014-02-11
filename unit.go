package antfarm

import "fmt"

type Unit interface {
	Actor
	Mover
}

type Worm struct {
	RandomMover
	id    int
	delay int
}

func (this *Worm) setId(id int) {
	this.id = id
}

func (this Worm) busy() bool {
	return this.delay > 0
}

func (this Worm) rollInit() int {
	return d6() + 10
}

func (this *Worm) tic(world *World) {
	if this.delay > 0 {
		this.delay--
	} else {
		this.act(world)
	}
}

func (this *Worm) act(world *World) {
	baseDelay := this.Speed
	init := this.rollInit()
	this.delay = (baseDelay / init)
	this.move(world)
	fmt.Println("Worm", this.id, "is moving to", this.Location.X, this.Location.Y)
}

func makeWorm(w World) *Worm {
	worm := new(Worm)
	worm.Speed = 1000
	worm.Location = w.Random()
	worm.delay = 0
	return worm
}
