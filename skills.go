package af

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
	switch this.Stat {
	case "strength":
		skill.Stat = ST
	case "agility":
		skill.Stat = AG
	case "toughness":
		skill.Stat = TO
	case "endurance":
		skill.Stat = EN
	case "health":
		skill.Stat = HT
	case "reaction":
		skill.Stat = RE
	case "willpower":
		skill.Stat = WP
	case "logic":
		skill.Stat = LG
	case "Intuiton":
		skill.Stat = IN
	case "charisma":
		skill.Stat = CH
	}
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

func (this BasicSkill) getStat() stat {
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
