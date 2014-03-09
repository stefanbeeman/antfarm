package rpg

type Material interface {
	GetName() string
	GetHardness() int
	GetFracture() int
	GetMelt() int
	GetFlammable() bool
	GetValue() int
}

type BasicMaterial struct {
	Name      string
	Hardness  int
	Fracture  int
	Melt      int
	Flammable bool
	Value     int
}

func (this BasicMaterial) GetName() string    { return this.Name }
func (this BasicMaterial) GetHardness() int   { return this.Hardness }
func (this BasicMaterial) GetFracture() int   { return this.Fracture }
func (this BasicMaterial) GetMelt() int       { return this.Melt }
func (this BasicMaterial) GetFlammable() bool { return this.Flammable }
func (this BasicMaterial) GetValue() int      { return this.Value }

var Materials map[string]Material
