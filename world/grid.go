package world

import (
	. "github.com/stefanbeeman/antfarm/common"
)

type Grid interface {
	width() int
	height() int
	contains(int, int) bool
	get(Location) Cell
	set(Location, Cell)
	All() [][]Cell
	Display() Display
}

type BasicGrid struct {
	W     int
	H     int
	Cells [][]Cell
}

func (this BasicGrid) width() int    { return this.W }
func (this BasicGrid) height() int   { return this.H }
func (g BasicGrid) size() (int, int) { return g.W, g.H }

func (g BasicGrid) contains(x, y int) bool {
	return (x >= 0) && (x < g.W) && (y >= 0) && (y < g.H)
}

func (g BasicGrid) get(l Location) Cell {
	x, y := l.Coords()
	return g.Cells[y][x]
}

func (g BasicGrid) set(l Location, c Cell) {
	x, y := l.Coords()
	g.Cells[y][x] = c
}

func (g BasicGrid) slice(x, y, w, h int) BasicGrid {
	sliced := g.Cells[y : y+h]
	for y := 0; y < h; y++ {
		sliced[y] = sliced[y][x : x+w]
	}
	return BasicGrid{w, h, sliced}
}

func (this BasicGrid) All() [][]Cell {
	return this.Cells
}

type DisplayBasicGrid struct {
	Cells [][]Display
}

func (this BasicGrid) Display() Display {
	cells := make([][]Display, this.H)
	for y, row := range this.Cells {
		cells[y] = make([]Display, this.W)
		for x, cell := range row {
			cells[y][x] = cell.Display()
		}
	}
	return DisplayBasicGrid{cells}
}

func MakeGrid(w int, h int) BasicGrid {
	g := BasicGrid{w, h, make([][]Cell, h)}
	for y := 0; y < h; y++ {
		g.Cells[y] = make([]Cell, w)
		for x := 0; x < w; x++ {
			g.Cells[y][x] = new(BasicCell)
		}
	}
	return g
}
