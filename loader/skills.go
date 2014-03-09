package loader

import (
	"github.com/stefanbeeman/antfarm/rpg"
)

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

func loadSkills(root string) {
	rpg.Skills = make(map[string]rpg.Skill)
	skills := make(map[string]protoskill)
	path := root + "/skills.yml"
	loadFile(path, skills)
	for name, skill := range skills {
		skill.name = name
		rpg.Skills[name] = skill.build()
	}
}
