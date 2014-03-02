package ai

import (
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/storage"
)

type Memory interface {
  storage.WorldState
  MoveCost(Location) (int, bool)
}

type OmniscientMemory struct {
	world storage.WorldState
}

func (this *OmniscientMemory) GetCell(p Location) storage.Cell   { return this.world.GetCell(p) }
func (this *OmniscientMemory) Contains(p Location) bool          { return this.world.Contains(p) }
func (this *OmniscientMemory) Update(p Location, c storage.Cell) { return }
func (this *OmniscientMemory) GetAll() storage.Grid              { return storage.MakeGrid(1, 1) }
func (this *OmniscientMemory) MoveCost(l Location) (int, bool) {
  cell := this.GetCell(l)
  return 10, cell.GetSolid()
}