package loader

import (
	"github.com/stefanbeeman/antfarm/rpg"
)

func loadMaterials(root string) {
	rpg.Materials = make(map[string]rpg.Material)
	mats := make(map[string]rpg.BasicMaterial)
	path := root + "/materials.yml"
	loadFile(path, mats)
	for name, mat := range mats {
		mat.Name = name
		rpg.Materials[name] = mat
	}
}
