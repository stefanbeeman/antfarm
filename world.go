package antfarm

type World struct {
	Grid2D
	Now       int
	Materials []Material
	Actors    []Actor
}

func (this *World) tic() {
	this.Now++
	for _, actor := range this.Actors {
		actor.tic(this)
	}
}

func (this World) Run(tics int) {
	for i := 0; i < tics; i++ {
		this.tic()
	}
}

func (this *World) addActor(a Actor) {
	this.Actors = append(this.Actors, a)
}

func MakeWorld(data string, width int, height int, worms int) World {
	g := makeGrid2D(width, height)
	m := LoadMaterials(data)
	u := make([]Actor, 0)
	w := World{g, 0, m, u}

	for y, row := range w.Cells {
		for x, _ := range row {
			p := Point{x, y}
			c := makeCell(p)
			if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
				c.set("solid", 1)
			}
			w.set(p, c)
		}
	}

	for n := 0; n < worms; n++ {
		rp := w.random()
		a := makeWorm(rp)
		w.addActor(a)
	}
	return w
}
