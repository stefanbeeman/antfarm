package antfarm

import (
	"github.com/soundcloud/goyaml"
	. "github.com/stefanbeeman/antfarm/common"
	"github.com/stefanbeeman/antfarm/rpg"
	"io/ioutil"
)

type Loader interface {
	setRoot(string)
	load(string, interface{})
	loadMaterials()
}

type YmlLoader struct {
	root string
}

func (this *YmlLoader) setRoot(path string) {
	this.root = path
}

func (this YmlLoader) load(path string, pointer interface{}) {
	Console.Meh("Loading " + path)
	buffer, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		Console.Broke(readErr)
	}
	marshErr := goyaml.Unmarshal(buffer, pointer)
	if marshErr != nil {
		Console.Broke(marshErr)
	}
}

func (this YmlLoader) loadMaterials() {
	files, _ := ioutil.ReadDir(this.root + "/materials")
	for _, file := range files {
		mat := new(rpg.BasicMaterial)
		path := this.root + "/materials/" + file.Name()
		this.load(path, mat)
		rpg.Materials[mat.Name] = mat
	}
}

// func (this YmlLoader) loadSkills() map[string]Skill {
// 	skills := make(map[string]Skill)
// 	files, _ := ioutil.ReadDir(this.root + "/skills")
// 	for _, file := range files {
// 		skill := new(ProtoBasicSkill)
// 		path := this.root + "/skills/" + file.Name()
// 		this.load(path, skill)
// 		skills[skill.Name] = skill.build()
// 	}
// 	return skills
// }

var yml = new(YmlLoader)
