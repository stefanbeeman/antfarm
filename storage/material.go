package storage

type Material interface {
	getName() string
	getStructure() int
	getHardness() int
}

type BasicMaterial struct {
	Name      string
	Color     string
	Structure int
	Hardness  int
}

func (this BasicMaterial) getName() string {
	return this.Name
}

func (this BasicMaterial) getStructure() int {
	return this.Structure
}

func (this BasicMaterial) getHardness() int {
	return this.Hardness
}
