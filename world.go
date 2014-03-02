package af

type WorldState interface {
  GetCell(Location) Cell
  Update(Location, Cell)
  Contains(Location) bool
}

type BasicWorld struct {
	grid BasicGrid
}

func (this BasicWorld) GetCell(l Location) Cell {
	return this.grid.get(l.coords)
}

func (this BasicWorld) Update(l Location, newCell Cell) {
	cell := this.grid.get(l.coords())
	cell = newCell
}

func (this BasicWorld) Contains(l Location) bool {
	return this.grid.contains(l.coords())
}

func makeWorld(w,h int) BasicWorld {
	return BasicWorld{
		makeGrid(w,h),
	}
}

func makeGrid(width int, height int) BasicGrid {
 g := Grid2D{w,h, make([][]Cell, height)}
 for y := 0; y < height; y++ {
   g.Cells[y] = make([]Cell, width)
   for x := 0; x < width; x++ {
     g.Cells[y][x] = new(BasicCell)
   }
 }
 return g
}