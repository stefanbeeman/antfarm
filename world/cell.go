package world

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/rpg"
)

type Cell interface {
	where() Location
	getMat() rpg.Material
	setMat(rpg.Material)
	GetSolid() bool
	SetSolid(bool)
	getData(string) int
	setData(string, int)
	Display() Display
}

type BasicCell struct {
	location Location
	Material rpg.Material
	Solid    bool
	Data     map[string]int
}

func (this BasicCell) where() Location {
	return this.location
}

func (this BasicCell) getMat() rpg.Material {
	return this.Material
}

func (this *BasicCell) setMat(mat rpg.Material) {
	this.Material = mat
}

func (this BasicCell) GetSolid() bool {
	return this.Solid
}

func (this *BasicCell) SetSolid(state bool) {
	this.Solid = state
}

func (this BasicCell) getData(prop string) int {
	return this.Data[prop]
}

func (this *BasicCell) setData(prop string, value int) {
	this.Data[prop] = value
}

func MakeCell(p Point, mat string, solid bool) Cell {
	c := BasicCell{
		p,
		rpg.Materials[mat],
		solid,
		make(map[string]int),
	}
	return &c
}

type DisplayBasicCell struct {
	Solid    bool
	Material string
}

func (this BasicCell) Display() Display {
	solid := this.Solid
	mat := this.Material.GetName()
	return DisplayBasicCell{solid, mat}
}
