package antfarm

import "math/rand"

type Mover interface {
	move(*Unit, Point)
	moveTo(*Unit, Point) Task
}

type RandomWalker interface{}

func (this RandomWalker) move(u *Unit, vector Point) {
	u
}

func (this RandomWalker) moveTo(u *Unit, p Point) Task {
	return BasicTask{
		func() bool {
			return u.Position.equals(p)
		},
		func() Action {
			i := rand.Intn(4)
			switch i {
			case 0:
				return func() {
					u.move(NORTH)
				}
			case 1:
				return func() {
					u.move(SOUTH)
				}
			case 2:
				return func() {
					u.move(EAST)
				}
			case 3:
				return func() {
					u.move(WEST)
				}
			}
		},
	}
}

type Speed struct {
	Walk   int
	Run    int
	Sprint int
}

func (this *Unit) makeWalkAction(direction Point) {}

func (this *Unit) makeRunAction(direction Point) {}
