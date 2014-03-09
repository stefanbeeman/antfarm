package rpg

var Skills map[string]Skill

type Skill interface {
	getName() string
	getTitle() string
	getStats() []int
	getDefaults() map[string]int
}

type BasicSkill struct {
	Name     string
	Title    string
	Stats    []int
	Defaults map[string]int
}

func (this BasicSkill) getName() string             { return this.Name }
func (this BasicSkill) getTitle() string            { return this.Title }
func (this BasicSkill) getStats() []int             { return this.Stats }
func (this BasicSkill) getDefaults() map[string]int { return this.Defaults }

type SkillLevel interface {
	get() (int, int)
	getTN() int
	getSkill() Skill
	getBase() int
	setBase(int)
	getShade() int
	setShade(int)
	getMod(string) int
	setMod(string, int)
	clearMod(string)
	resetMods(string)
	getApt() int
	setApt(int)
	getXP() int
	awardXP(int)
}

type BasicSkillLevel struct {
	Skill Skill
	Base  int
	Shade int
	Mods  map[string]int
	Apt   int
	XP    int
}

func (this BasicSkillLevel) getSkill() Skill { return this.Skill }

func (this BasicSkillLevel) getBase() int       { return this.Base }
func (this *BasicSkillLevel) setBase(value int) { this.Base = value }

func (this BasicSkillLevel) getShade() int       { return this.Shade }
func (this *BasicSkillLevel) setShade(value int) { this.Shade = value }

func (this BasicSkillLevel) getMod(mod string) int            { return this.Mods[mod] }
func (this *BasicSkillLevel) setMod(mod string, value int)    { this.Mods[mod] = value }
func (this *BasicSkillLevel) clearMod(mod string)             { this.Mods[mod] = 0 }
func (this *BasicSkillLevel) resetMods(mod string, value int) { this.Mods = make(map[string]int) }

func (this BasicSkillLevel) getApt() int       { return this.Apt }
func (this *BasicSkillLevel) setApt(value int) { this.Apt = value }

func (this BasicSkillLevel) getXP() int { return this.XP }
func (this *BasicSkillLevel) awardXP(value int) {
	this.XP += value
	if this.Base < GRANDMASTER {
		next := this.Base + 1
		if this.XP >= (next * SKILLXP) {
			this.Base += 1
			this.XP = 0
		}
	}
}

type Skilled interface {
	Statted
	GetSkill(string) (int, int)
	SetSkillMod(string, string, int)
	ClearSkillMod(string, string)
	ResetSkillMods(string)
	GetSkillApt(string)
	AwardSkillXP(string, int)
	GetSkillStats(string) []int
	RollSkill(string) Roll
	TestSkill(string, int) (int, bool)
}

type BasicSkilled struct {
	BasicStatted
	Skills map[string]SkillLevel
}

func (this BasicSkilled) dispatch(which string) SkillLevel { return this.Skills[which] }

func (this BasicSkilled) GetSkill(which string) (int, int) {
	result := this.dispatch(which)
	return result.getBase(), result.getShade()
}

func (this BasicSkilled) SetSkillMod(which string, mod string, value int) {
	this.dispatch(which).setMod(mod, value)
}

func (this *BasicSkilled) ClearSkillMod(which string, mod string) {
	this.dispatch(which).clearMod(mod)
}

func (this *BasicSkilled) ResetSkillMods(which string, mod string) {
	this.dispatch(which).resetMods(mod)
}

func (this *BasicSkilled) AwardSkillXP(which string, award int) {
	this.dispatch(which).awardXP(award)
}

func (this BasicSkilled) GetSkillApt(which string) int { return this.dispatch(which).getApt() }

func (this BasicSkilled) GetSkillStats(which string) []int {
	return this.dispatch(which).getSkill().getStats()
}

func (this BasicSkilled) RollSkill(which string) Roll {
	level, shade := this.GetSkill(which)
	tn := 10 - level
	dice := 0
	stats := this.GetSkillStats(which)
	for _, stat := range stats {
		statDice, statShade := this.GetStat(stat)
		dice += statDice
		if statShade > shade {
			shade = statShade
		}
	}
	dice = dice / len(stats)
	return BasicRoll{dice, tn, shade}
}

func (this *BasicSkilled) TestSkill(which string, target int) (int, bool) {
	roll := this.RollSkill(which)
	hits, botch := Dice.RollTest(roll, target)
	apt := this.GetSkillApt(which)
	if botch {
		this.AwardSkillXP(which, 1)
	} else if hits < apt {
		this.AwardSkillXP(which, 1)
	}
	return hits, botch
}
