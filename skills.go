package antfarm

type Skill interface {
	getName() string
	getStat() int
	getDefaults() map[string]int
	getTitle() string
	getDesc() string
}

type ProtoBasicSkill struct {
	Name     string
	Stat     string
	Defaults map[string]int
	Title    string
	Desc     string
}

func (this ProtoBasicSkill) build() Skill {
	skill := new(BasicSkill)
	skill.Name = this.Name
	skill.Stat = 0
	skill.Defaults = this.Defaults
	skill.Title = this.Title
	skill.Desc = this.Desc
	return skill
}

type BasicSkill struct {
	Name     string
	Stat     int
	Defaults map[string]int
	Title    string
	Desc     string
}

func (this BasicSkill) getName() string {
	return this.Name
}

func (this BasicSkill) getStat() int {
	return this.Stat
}

func (this BasicSkill) getDefaults() map[string]int {
	return this.Defaults
}

func (this BasicSkill) getTitle() string {
	return this.Title
}

func (this BasicSkill) getDesc() string {
	return this.Desc
}
