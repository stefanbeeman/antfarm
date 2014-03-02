package common

type Location interface {
	At(Location) bool
	DistanceTo(Location) int
	AsPoint() Point
	Neighbors() []Location
	Coords() (int, int)
}

type MutableLocation interface {
	Location
	SetPosition(Location)
}

func (this *Point) SetPosition(that Location) { this.X, this.Y = that.Coords() }

type Point struct {
	X int
	Y int
}

func (this Point) Coords() (int, int)    { return this.X, this.Y }
func (this Point) At(that Location) bool { return this.AsPoint() == that.AsPoint() }
func (this Point) AsPoint() Point {
	x, y := this.Coords()
	return Point{x, y}
}

func (this Point) DistanceTo(that Location) int {
	x1, y1 := this.Coords()
	x2, y2 := that.Coords()
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
	x, y := this.Coords()
	return []Location{
		Point{x + 1, y},
		Point{x, y + 1},
		Point{x - 1, y},
		Point{x, y - 1},
	}
}
