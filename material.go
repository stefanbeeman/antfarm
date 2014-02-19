package antfarm

import (
	"fmt"
	"github.com/soundcloud/goyaml"
	"io/ioutil"
)

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

func LoadMaterials(data string) []Material {
	mats := make([]Material, 0)
	files, _ := ioutil.ReadDir(data + "/materials")
	for _, file := range files {
		mat := new(BasicMaterial)
		buffer, _ := ioutil.ReadFile(data + "/materials/" + file.Name())
		_ = goyaml.Unmarshal(buffer, mat)
		fmt.Println(mat)
		mats = append(mats, mat)
	}
	fmt.Println(mats[0])
	return mats
}
