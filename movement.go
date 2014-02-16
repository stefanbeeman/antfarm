package antfarm

type Speed struct {
	Walk   int
	Run    int
	Sprint int
}

func (this *Unit) makeWalkAction(direction Point) {}

func (this *Unit) makeRunAction(direction Point) {}
