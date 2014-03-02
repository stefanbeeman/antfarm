package af

import (
	"strconv"
	"time"
)

type WorldState interface {
  GetCell(Location) Cell
  Update(Location, Cell)
  Contains(Location) bool
}

type BasicWorld struct {
	grid Grid2D
}

func (this BasicWorld) GetCell(l Location) Cell {
	return this.grid.get(l.coords)
}

func (this BasicWorld) Update(l Location, newCell Cell) {
	cell := this.grid.get(l.coords())
	cell = newCell
}

func (this BasicWorld) Contains(l Location) {
	return this.grid.contains(l.coords())
}