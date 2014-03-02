package af

type WorldState interface {
	GetCell(Location) Cell
	GetAll() Grid
	Update(Location, Cell)
	Contains(Location) bool
}

type BasicWorld struct {
	Grid BasicGrid
}

func (this BasicWorld) GetCell(l Location) Cell {
	x, y := l.coords()
	return this.Grid.get(x, y)
}

func (this BasicWorld) GetAll() Grid {
	return this.Grid
}

func (this BasicWorld) Update(l Location, newCell Cell) {
	x, y := l.coords()
	this.Grid.set(x, y, newCell)
}

func (this BasicWorld) Contains(l Location) bool {
	return this.Grid.contains(l.coords())
}

func makeWorld(w, h int) BasicWorld {
	return BasicWorld{
		makeGrid(w, h),
	}
}

func makeGrid(w int, h int) BasicGrid {
	g := BasicGrid{w, h, make([][]Cell, h)}
	for y := 0; y < h; y++ {
		g.Cells[y] = make([]Cell, w)
		for x := 0; x < w; x++ {
			g.Cells[y][x] = new(BasicCell)
		}
	}
	return g
}
