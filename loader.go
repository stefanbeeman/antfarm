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
	rpg.Materials = make(map[string]rpg.Material)
	files, _ := ioutil.ReadDir(this.root + "/materials")
	for _, file := range files {
		mat := new(rpg.BasicMaterial)
		path := this.root + "/materials/" + file.Name()
		this.load(path, mat)
		rpg.Materials[mat.Name] = mat
	}
}

type protoskill struct {
	name     string
	title    string
	stats    []string
	defaults map[string]int
}

func (this protoskill) build() rpg.Skill {
	result := new(rpg.BasicSkill)
	result.Name = this.name
	result.Title = this.title
	result.Stats = make([]int, len(this.stats))
	for i, statstring := range this.stats {
		stat := rpg.ParseStat(statstring)
		result.Stats[i] = stat
	}
	result.Defaults = this.defaults
	return result
}

func (this YmlLoader) loadSkills() {
	rpg.Skills = make(map[string]rpg.Skill)
	files, _ := ioutil.ReadDir(this.root + "/skills")
	for _, file := range files {
		skill := new(protoskill)
		path := this.root + "/skills/" + file.Name()
		this.load(path, skill)
		rpg.Skills[skill.name] = skill.build()
	}
}

var yml = new(YmlLoader)
