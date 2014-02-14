package antfarm

type Mover struct {
	Position Point
}

func (this *Mover) Move(vector Point) {
	this.Position.add(vector)
}
