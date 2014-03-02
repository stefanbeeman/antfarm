package common

type Location interface {
  At(Location) bool
  DistanceTo(Location) int
  AsPoint() Point
  Neighbors() []Location
  coords() (int, int)
}

type MutableLocation interface {
  Location
  SetPosition(Location)
}

func (this *Point) SetPosition(that Location) { this.x, this.y = that.coords() }

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