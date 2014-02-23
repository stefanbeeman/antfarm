package af

type Material interface {
	getName() string
	getGlyph() string
	getStructure() int
	getHardness() int
}

type BasicMaterial struct {
	Name      string
	Glyph     string
	Structure int
	Hardness  int
}

func (this BasicMaterial) getName() string {
	return this.Name
}

func (this BasicMaterial) getGlyph() string {
	return this.Glyph
}

func (this BasicMaterial) getStructure() int {
	return this.Structure
}

func (this BasicMaterial) getHardness() int {
	return this.Hardness
}
