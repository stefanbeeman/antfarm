package rpg

type Material interface {
	GetName() string
	GetStructure() int
	GetHardness() int
}

type BasicMaterial struct {
	Name      string
	Structure int
	Hardness  int
}

func (this BasicMaterial) GetName() string {
	return this.Name
}

func (this BasicMaterial) GetStructure() int {
	return this.Structure
}

func (this BasicMaterial) GetHardness() int {
	return this.Hardness
}

var Materials map[string]Material
