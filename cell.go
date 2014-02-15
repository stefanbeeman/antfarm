package antfarm

import "strconv"

type Cell struct {
	location Point
	Data     map[string]int
}

func (c Cell) place(p Point) Cell {
	c.location = p
	return c
}

func (c Cell) where() Point {
	return c.location
}

func (c Cell) get(prop string) int {
	return c.Data[prop]
}

func (c Cell) set(prop string, value int) {
	c.Data[prop] = value
}

func (c Cell) show() string {
	return "#"
}

func (c Cell) showData(prop string) string {
	data := c.get(prop)
	return strconv.Itoa(data)
}

func makeCell(p Point) Cell {
	c := Cell{
		p,
		make(map[string]int),
	}
	c.set("material", 0)
	c.set("solid", 0)
	return c
}
