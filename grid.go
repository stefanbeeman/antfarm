package af

type Location interface {
	At(Location) bool
	DistanceTo(Location) int
	AsPoint() Point
	SetPosition(Location)
	Neighbors() []Location
	coords() (int, int)
}

type Point struct {
	x int
	y int
}

func (this Point) coords() (int, int)    { return this.x, this.y }
func (this Point) At(that Location) bool { return this.AsPoint() == that.AsPoint() }
func (this Point) AsPoint() Point {
	x, y := this.coords()
	return Point{x, y}
}
func (this Point) SetPosition(that Location) { this.x, this.y = that.coords() }

func (this Point) DistanceTo(that Location) int {
	x1, y1 := this.coords()
	x2, y2 := that.coords()
	dx, dy := x1-x2, y1-y2
	if dx < 0 {
		dx = -1 * dx
	}
	if dy < 0 {
		dy = -1 * dy
	}
	return dx + dy
}

func (this Point) Neighbors() []Location {
	x, y := this.coords()
	return []Location{
		Point{x + 1, y},
		Point{x, y + 1},
		Point{x - 1, y},
		Point{x, y - 1},
	}
}

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
