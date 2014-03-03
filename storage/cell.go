package storage

import (
	. "github.com/stefanbeeman/antfarm/common"
)

type Cell interface {
	where() Location
	getMat() Material
	setMat(Material)
	GetSolid() bool
	SetSolid(bool)
	getData(string) int
	setData(string, int)
	Display() Display
}

type BasicCell struct {
	location Location
	Material Material
	Solid    bool
	Data     map[string]int
}

func (this BasicCell) where() Location {
	return this.location
}

func (this BasicCell) getMat() Material {
	return this.Material
}

func (this *BasicCell) setMat(mat Material) {
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

func MakeCell(p Point, mat Material, solid bool) Cell {
	c := BasicCell{
		p,
		mat,
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
	mat := this.Material.getName()
	return DisplayBasicCell{solid, mat}
}
