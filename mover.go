package antfarm

var (
	NORTH = Point{0, -1}
	SOUTH = Point{0, 1}
	EAST  = Point{1, 0}
	WEST  = Point{-1, 0}
)

type Mover struct {
	Position Point
}

func (this *Mover) Move(vector Point) {
	this.Position.Add(vector)
}
