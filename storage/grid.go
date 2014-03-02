package storage

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
