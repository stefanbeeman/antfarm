package antfarm

import (
	"fmt"
	"strconv"
)

type AFCell struct {
	location Point
	Data     map[string]int
}

func (c AFCell) Place(p Point) Cell {
	c.location = p
	return c
}

func (c AFCell) Where() Point {
	return c.location
}

func (c AFCell) Get(prop string) int {
	return c.Data[prop]
}

func (c AFCell) Set(prop string, value int) {
	c.Data[prop] = value
}

func (c AFCell) Show() string {
	return "[ ]"
}

func (c AFCell) ShowData(prop string) string {
	data := c.Get(prop)
	return "[" + strconv.Itoa(data) + "]"
}

func MakeCell(p Point) AFCell {
	c := AFCell{p, make(map[string]int)}
	c.Set("material", 0)
	c.Set("solid", 0)
	return c
}

type World struct {
	Grid2D
	Now       int
	Materials []Material
	Units     []Unit
}

func (w World) Step() {
	w.Now++
	for _, unit := range w.Units {
		unit.Act(&w)
	}
}

func (w World) Resolve() {
	//fmt.Println("Resolve")
}

func (w World) Run(tics int) {
	for tic := 0; tic < tics; tic++ {
		w.Step()
		w.Resolve()
	}
}

func (w *World) AddUnit(u Unit) {
	w.Units = append(w.Units, u)
}

func MakeWorld(data string, width int, height int, worms int) World {
	g := MakeGrid2D(width, height)
	m := LoadMaterials(data)
	u := make([]Unit, 0)
	w := World{g, 0, m, u}

	for y, row := range w.Cells {
		for x, _ := range row {
			p := Point{x, y}
			c := MakeCell(p)
			if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
				c.Set("solid", 1)
			}
			w.Set(p, c)
		}
	}

	for n := 0; n < worms; n++ {
		p := w.Random()
		worm := Worm{p, 0}
		w.AddUnit(worm)
	}

	return w
}
