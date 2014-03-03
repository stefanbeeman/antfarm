package storage

import (
	. "github.com/stefanbeeman/antfarm/common"
)

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
	return this.Grid.get(l)
}

func (this BasicWorld) GetAll() Grid {
	return this.Grid
}

func (this BasicWorld) Update(l Location, newCell Cell) {
	this.Grid.set(l, newCell)
}

func (this BasicWorld) Contains(l Location) bool {
	return this.Grid.contains(l.Coords())
}

func MakeWorld(w, h int) BasicWorld {
	return BasicWorld{
		MakeGrid(w, h),
	}
}
