package antfarm

import "math/rand"

var (
	NORTH   = Point{0, 1}
	SOUTH   = Point{0, -1}
	EAST    = Point{1, 0}
	WEST    = Point{-1, 0}
	NOWHERE = Point{0, 0}
)

type Mover interface {
	move(*Unit, Point)
	moveTo(*Unit, Point) Task
}

type RandomWalker struct {
	Speed int
}

func (this RandomWalker) move(u *BasicUnit, vector Point) {
	u.Position = u.Position.add(vector)
}

func (this RandomWalker) moveTo(u *BasicUnit, p Point) Task {
	return BasicTask{
		func() bool {
			return u.Position.equals(p)
		},
		func() Action {
			fn := func() {}
			i := rand.Intn(4)
			switch i {
			case 0:
				fn = func() {
					this.move(u, NORTH)
				}
			case 1:
				fn = func() {
					this.move(u, SOUTH)
				}
			case 2:
				fn = func() {
					this.move(u, EAST)
				}
			case 3:
				fn = func() {
					this.move(u, WEST)
				}
			}
			act := BasicAction{
				this.Speed,
				fn,
			}
			return &act
		},
	}
}
