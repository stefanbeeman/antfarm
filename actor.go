package antfarm

import (
	"fmt"
)

type Actor interface {
	setId(int)
	act(i *World)
	busy() bool
	rollInit() int
	tic(*World)
}

type RootActor struct {
	id    int
	delay int
}

func (this RootActor) setId(id int) {
	this.id = id
}

func (this *RootActor) act(world *World) {
	baseDelay := 1000
	init := this.rollInit()
	this.delay = (baseDelay / init)
	fmt.Println("Actor", this.id, " taking an action.")
}

func (this RootActor) busy() bool {
	return this.delay > 0
}

func (this RootActor) rollInit() int {
	return d6() + 10
}

func (this *RootActor) tic(world *World) {
	if this.delay > 0 {
		this.delay--
	} else {
		this.act(world)
	}
}
