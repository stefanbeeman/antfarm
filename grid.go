package antfarm

import (
	"fmt"
	"math"
	"math/rand"
)

type Cell interface {
	Place(Point) Cell
	Where() Point
	Get(string) int
	Set(string, int)
	Show() string
	ShowData(string) string
}

//Type signatures for common lambda functions used when iterating over grids.
type iterator func(Cell, Point)
type locator func(Cell, Point) bool

type Grid interface {
	Size() []int
	Contains(Point) bool
	Get(Point) Cell
	Set(Point, Cell)
	Each(iterator)
	Find(locator) []Point
	Slice(Point, Point) Grid
	Around(Point, int) Grid
	Random() Point
	RandomCurved(int) Point
	Show()
	ShowData(string)
}

type Point struct {
	X int
	Y int
}

var (
	NORTH = Point{0, -1}
	SOUTH = Point{0, 1}
	EAST  = Point{1, 0}
	WEST  = Point{-1, 0}
)

func (this Point) Add(that Point) Point {
	return Point{this.X + that.X, this.Y + that.Y}
}

func (here Point) DistanceTo(there Point) float64 {
	x := math.Pow(float64(there.X)-float64(here.X), 2)
	y := math.Pow(float64(there.Y)-float64(here.Y), 2)
	return math.Sqrt(x + y)
}

type Grid2D struct {
	Cells [][]Cell
}

func (this Grid2D) Width() int {
	return len(this.Cells[0])
}

func (g Grid2D) Height() int {
	return len(g.Cells)
}

func (g Grid2D) Size() []int {
	size := make([]int, 2)
	size[0] = g.Width()
	size[1] = g.Height()
	return size
}

func (g Grid2D) Contains(here Point) bool {
	return (here.X >= 0) && (here.X < g.Width()) && (here.Y >= 0) && (here.Y < g.Height())
}

func (g Grid2D) Get(here Point) Cell {
	return g.Cells[here.Y][here.X]
}

func (g Grid2D) Set(here Point, c Cell) {
	g.Cells[here.Y][here.X] = c
}

func (g Grid2D) Each(fn iterator) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			p := Point{x, y}
			c := g.Get(p)
			fn(c, p)
		}
	}
}

func (g Grid2D) Find(fn locator) []Point {
	result := make([]Point, 0)
	g.Each(func(c Cell, p Point) {
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

func (g Grid2D) Slice(start Point, end Point) Grid {
	lowX, highX := orderSlice(start.X, end.X, g.Width())
	lowY, highY := orderSlice(start.Y, end.Y, g.Height())

	sliced := g.Cells[lowY : highY+1]
	for y := 0; y < len(sliced); y++ {
		sliced[y] = sliced[y][lowX : highX+1]
	}
	return Grid2D{sliced}
}

func (g Grid2D) Around(center Point, r int) Grid {
	start := Point{center.X - r, center.Y - r}
	end := Point{center.X + r, center.Y + r}
	return g.Slice(start, end)
}

func (g Grid2D) Random() Point {
	x := rand.Intn(g.Width())
	y := rand.Intn(g.Height())
	return Point{x, y}
}

func (g Grid2D) RandomCurved(curve int) Point {
	x := 0
	y := 0
	for i := 0; i < curve; i++ {
		x = x + rand.Intn(g.Width())
		y = y + rand.Intn(g.Height())
	}
	x = int(float64(x) / float64(curve))
	y = int(float64(y) / float64(curve))
	return Point{x, y}
}

func (g Grid2D) Show() {
	for y := 0; y < g.Height(); y++ {
		line := ""
		for x := 0; x < g.Width(); x++ {
			c := g.Get(Point{x, y})
			line = line + "[" + c.Show() + "]"
			if x < (g.Width() - 1) {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func (g Grid2D) ShowData(layer string) {
	for y := 0; y < g.Height(); y++ {
		line := ""
		for x := 0; x < g.Width(); x++ {
			c := g.Get(Point{x, y})
			line = line + "[" + c.ShowData(layer) + "]"
			if x < (g.Width() - 1) {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func MakeGrid2D(width int, height int) Grid2D {
	g := Grid2D{make([][]Cell, height)}
	for y := 0; y < height; y++ {
		g.Cells[y] = make([]Cell, width)
	}
	return g
}
