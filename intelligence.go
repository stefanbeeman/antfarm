package antfarm

type Thinker interface {
	think(*Unit) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *Unit) Task {
	return u.mover.MoveTo(u, Point{0, 0})
}
