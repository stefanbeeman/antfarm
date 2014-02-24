package af

type Material interface {
	getName() string
	getGlyph() string
	getColor() string
	getStructure() int
	getHardness() int
}

type BasicMaterial struct {
	Name      string
	Glyph     string
	Color     string
	Structure int
	Hardness  int
}

func (this BasicMaterial) getName() string {
	return this.Name
}

func (this BasicMaterial) getGlyph() string {
	return this.Glyph
}

func (this BasicMaterial) getColor() string {
	return this.Color
}

func (this BasicMaterial) getStructure() int {
	return this.Structure
}

func (this BasicMaterial) getHardness() int {
	return this.Hardness
}
