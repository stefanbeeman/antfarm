package af

type OmniscientMemory struct {
  world *BasicWorld
}

func (this OmniscientMemory) GetCell(p Location) Cell { return this.world.GetCell(p)}
func (this OmniscientMemory) Contains(p Location) { return this.world.Contains(p) }
func (this OmniscientMemory) Update(p Location, c Cell) { return }