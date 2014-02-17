package antfarm

type Thinker interface {
	think(*BasicUnit, *World) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *BasicUnit, w *World) Task {
	return u.mover.moveTo(u, Point{0, 0})
}
