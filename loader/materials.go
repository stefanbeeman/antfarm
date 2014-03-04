package loader

import (
	"github.com/stefanbeeman/antfarm/rpg"
	"io/ioutil"
)

func loadMaterials(root string) {
	rpg.Materials = make(map[string]rpg.Material)
	files, _ := ioutil.ReadDir(root + "/materials")
	for _, file := range files {
		mat := new(rpg.BasicMaterial)
		path := root + "/materials/" + file.Name()
		loadFile(path, mat)
		rpg.Materials[mat.Name] = mat
	}
}
