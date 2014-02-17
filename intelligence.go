package antfarm

type Thinker interface {
	think(*BasicUnit) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *BasicUnit) Task {
	return u.mover.moveTo(u, Point{0, 0})
}
