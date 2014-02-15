package antfarm

import (
	"fmt"
	"math"
	"math/rand"
)

//Type signatures for common lambda functions used when iterating over grids.
type iterator func(Cell, Point)
type locator func(Cell, Point) bool

type Point struct {
	X int
	Y int
}

func (this Point) add(that Point) Point {
	return Point{this.X + that.X, this.Y + that.Y}
}

func (here Point) distanceTo(there Point) float64 {
	x := math.Pow(float64(there.X)-float64(here.X), 2)
	y := math.Pow(float64(there.Y)-float64(here.Y), 2)
	return math.Sqrt(x + y)
}

func (this Point) vectorTo(that Point) Point {
	X := 0
	Y := 0
	if this.X < that.X {
		X = 1
	} else if this.X > that.X {
		X = -1
	}
	if this.Y < that.Y {
		Y = 1
	} else if this.Y < that.Y {
		Y = -1
	}
	return Point{X, Y}
}

type Grid2D struct {
	Cells [][]Cell
}

func (this Grid2D) width() int {
	return len(this.Cells[0])
}

func (g Grid2D) height() int {
	return len(g.Cells)
}

func (g Grid2D) size() []int {
	size := make([]int, 2)
	size[0] = g.width()
	size[1] = g.height()
	return size
}

func (g Grid2D) contains(here Point) bool {
	return (here.X >= 0) && (here.X < g.width()) && (here.Y >= 0) && (here.Y < g.height())
}

func (g Grid2D) get(here Point) Cell {
	return g.Cells[here.Y][here.X]
}

func (g Grid2D) set(here Point, c Cell) {
	g.Cells[here.Y][here.X] = c
}

func (g Grid2D) each(fn iterator) {
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			p := Point{x, y}
			c := g.get(p)
			fn(c, p)
		}
	}
}

func (g Grid2D) find(fn locator) []Point {
	result := make([]Point, 0)
	g.each(func(c Cell, p Point) {
		if fn(c, p) {
			result = append(result, p)
		}
	})
	return result
}

func orderSlice(a int, b int, limit int) (low int, high int) {
	l := 0
	h := 0
	if a < b {
		l, h = a, b
	} else {
		l, h = b, a
	}
	if l < 0 {
		l = 0
	}
	if h > (limit - 1) {
		h = (limit - 1)
	}
	return l, h
}

func (g Grid2D) slice(start Point, end Point) Grid2D {
	lowX, highX := orderSlice(start.X, end.X, g.width())
	lowY, highY := orderSlice(start.Y, end.Y, g.height())

	sliced := g.Cells[lowY : highY+1]
	for y := 0; y < len(sliced); y++ {
		sliced[y] = sliced[y][lowX : highX+1]
	}
	return Grid2D{sliced}
}

func (g Grid2D) around(center Point, r int) Grid2D {
	start := Point{center.X - r, center.Y - r}
	end := Point{center.X + r, center.Y + r}
	return g.slice(start, end)
}

func (g Grid2D) random() Point {
	x := rand.Intn(g.width())
	y := rand.Intn(g.height())
	return Point{x, y}
}

func (g Grid2D) randomCurved(curve int) Point {
	x := 0
	y := 0
	for i := 0; i < curve; i++ {
		x = x + rand.Intn(g.width())
		y = y + rand.Intn(g.height())
	}
	x = x / curve
	y = y / curve
	return Point{x, y}
}

func (g Grid2D) show() {
	for y := 0; y < g.height(); y++ {
		line := ""
		for x := 0; x < g.width(); x++ {
			c := g.get(Point{x, y})
			line = line + "[" + c.show() + "]"
			if x < (g.width() - 1) {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func (g Grid2D) showData(layer string) {
	for y := 0; y < g.height(); y++ {
		line := ""
		for x := 0; x < g.width(); x++ {
			c := g.get(Point{x, y})
			line = line + "[" + c.showData(layer) + "]"
			if x < (g.width() - 1) {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func makeGrid2D(width int, height int) Grid2D {
	g := Grid2D{make([][]Cell, height)}
	for y := 0; y < height; y++ {
		g.Cells[y] = make([]Cell, width)
		for x := 0; x < width; x++ {
			l := Point{x, y}
			g.Cells[y][x] = makeCell(l)
		}
	}
	return g
}
