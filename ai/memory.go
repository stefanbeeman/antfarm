package ai

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/world"
)

type Memory interface {
	world.WorldState
	MoveCost(Location) (int, bool)
}

type OmniscientMemory struct {
	world world.WorldState
}

func (this *OmniscientMemory) GetCell(p Location) world.Cell   { return this.world.GetCell(p) }
func (this *OmniscientMemory) Contains(p Location) bool        { return this.world.Contains(p) }
func (this *OmniscientMemory) Update(p Location, c world.Cell) { return }
func (this *OmniscientMemory) GetAll() world.Grid              { return world.MakeGrid(1, 1) }
func (this *OmniscientMemory) MoveCost(l Location) (int, bool) {
	cell := this.GetCell(l)
	return 10, cell.GetSolid()
}
