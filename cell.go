package af

import (
	"fmt"
)

type Cell interface {
	where() Point
	getMat() Material
	setMat(Material)
	getSolid() bool
	setSolid(bool)
	getData(string) int
	setData(string, int)
	show()
	showData(string)
}

type BasicCell struct {
	location Point
	Material Material
	Solid    bool
	Data     map[string]int
}

func (this BasicCell) where() Point {
	return this.location
}

func (this BasicCell) getMat() Material {
	return this.Material
}

func (this BasicCell) setMat(mat Material) {
	this.Material = mat
}

func (this BasicCell) getSolid() bool {
	return this.Solid
}

func (this BasicCell) setSolid(state bool) {
	this.Solid = state
}

func (this BasicCell) getData(prop string) int {
	return this.Data[prop]
}

func (this BasicCell) setData(prop string, value int) {
	this.Data[prop] = value
}

func (this BasicCell) show() {
	fmt.Print(this.Material.getGlyph())
}

func (this BasicCell) showData(prop string) {
	fmt.Print(this.getData(prop))
}

func makeCell(p Point, mat Material, solid bool) Cell {
	c := BasicCell{
		p,
		mat,
		solid,
		make(map[string]int),
	}
	return &c
}
