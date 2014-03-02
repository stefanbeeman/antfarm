package af

import (
	"fmt"
	"math"
	"math/rand"
)

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

func (this Point) coords() (int, int) { return (x, y) }
func (this Point) At(that Location) bool { return this.coords() == that.coords() }
func (this Point) AsPoint() Point { return Point{this.coords()} }
func (this Point) SetPosition(that Location) { this.x, this.y = that.coords() }

func (this Point) DistanceTo(that Location) bool {
	x1, y1 := this.coords()
	x2, y2 := that.coords()
	return math.Abs(x1 - x2) + math.Abs(y1 - y2)
}

func (this Point) Neighbors() []Location {
	x, y := this.coords
	return []Location {
		Point{x+1, y},
		Point{x, y+1},
		Point{x-1, y},
		Point{x, y-1}
	}
}

type Grid interface {
	width() int
	height() int
	contains(int, int) bool
	get(int, int) Cell
}

type BasicGrid struct {
	width
	height
	Cells [][]Cell
}

func (this BasicGrid) width() int { return this.width }
func (this BasicGrid) height() int { return this.height }
func (g BasicGrid) size() (int, int) { return (g.width, g.height }

func (g BasicGrid) contains(x, y int) bool {
	return (x >= 0) && (x < g.width && (y >= 0) && (y < g.height
}

func (g BasicGrid) get(x, y int) Cell {
	return g.Cells[y][x]
}

func (g BasicGrid) slice(x, y, w, h int) BasicGrid {
	sliced := g.Cells[y : y+h]
	for y := 0; y < h; y++ {
		sliced[y] = sliced[y][x : x+w]
	}
	return BasicGrid{sliced}
}

func makeGrid2D(width int, height int) Grid2D {
 g := Grid2D{w,h, make([][]Cell, height)}
 for y := 0; y < height; y++ {
   g.Cells[y] = make([]Cell, width)
   for x := 0; x < width; x++ {
     g.Cells[y][x] = new(BasicCell)
   }
 }
 return g
}
