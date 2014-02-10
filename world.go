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

func MakeCell() AFCell {
	p := Point{0, 0}
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
		if tic%100 == 0 {
			fmt.Println("Current Tic: " + strconv.Itoa(tic))
		}
	}
}

func MakeWorld(width int, height int) World {
	c := MakeCell()
	g := MakeGrid2D(width, height, c)
	m := LoadMaterials()
	u := make([]Unit, 0)
	w := World{g, 0, m, u}
	w.Each(func(c Cell, p Point) {
		if p.X == 0 {
			c.Set("solid", 1)
			fmt.Println(c)
		}
	})
	return w
}
