package af

type Grid interface {
	width() int
	height() int
	contains(int, int) bool
	get(int, int) Cell
	set(int, int, Cell)
	all() [][]Cell
}

type BasicGrid struct {
	w     int
	h     int
	Cells [][]Cell
}

func (this BasicGrid) width() int    { return this.w }
func (this BasicGrid) height() int   { return this.h }
func (g BasicGrid) size() (int, int) { return g.w, g.h }

func (g BasicGrid) contains(x, y int) bool {
	return (x >= 0) && (x < g.w) && (y >= 0) && (y < g.h)
}

func (g BasicGrid) get(x, y int) Cell {
	return g.Cells[y][x]
}

func (g BasicGrid) set(x, y int, c Cell) {
	g.Cells[y][x] = c
}

func (g BasicGrid) slice(x, y, w, h int) BasicGrid {
	sliced := g.Cells[y : y+h]
	for y := 0; y < h; y++ {
		sliced[y] = sliced[y][x : x+w]
	}
	return BasicGrid{w, h, sliced}
}

func (this BasicGrid) all() [][]Cell {
	return this.Cells
}
