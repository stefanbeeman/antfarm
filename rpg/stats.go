package rpg

type StatLevel interface {
	get() (int, int)
	getBase() int
	setBase(int)
	getShade() int
	setShade(int)
	getMod(string) int
	setMod(string, int)
	clearMod(string)
	resetMods()
	getApt() int
	setApt(int)
	getMax() int
	setMax(int)
	getXP() int
	awardXP(int)
	roll() Roll
	test(int) (int, bool)
}

type BasicStatLevel struct {
	Base  int
	Shade int
	Mods  map[string]int
	Apt   int
	Max   int
	XP    int
}

func (this BasicStatLevel) get() (int, int) {
	value := this.Base
	for _, mod := range this.Mods {
		value += mod
	}
	return value, this.Shade
}

func (this BasicStatLevel) getBase() int       { return this.Base }
func (this *BasicStatLevel) setBase(value int) { this.Base = value }

func (this BasicStatLevel) getShade() int       { return this.Shade }
func (this *BasicStatLevel) setShade(value int) { this.Shade = value }

func (this BasicStatLevel) getMod(mod string) int         { return this.Mods[mod] }
func (this *BasicStatLevel) setMod(mod string, value int) { this.Mods[mod] = value }
func (this *BasicStatLevel) clearMod(mod string)          { this.Mods[mod] = 0 }
func (this *BasicStatLevel) resetMods()                   { this.Mods = make(map[string]int) }

func (this BasicStatLevel) getApt() int       { return this.Apt }
func (this *BasicStatLevel) setApt(value int) { this.Apt = value }

func (this BasicStatLevel) getMax() int       { return this.Max }
func (this *BasicStatLevel) setMax(value int) { this.Max = value }

func (this BasicStatLevel) getXP() int { return this.XP }
func (this *BasicStatLevel) awardXP(value int) {
	this.XP += value
	if this.Base < this.Max {
		next := this.Base + 1
		if this.XP >= (next * STATXP) {
			this.Base += 1
			this.XP = 0
		}
	}
}

func (this BasicStatLevel) roll() Roll {
	dice, shade := this.get()
	return BasicRoll{dice, 6, shade}
}

func (this *BasicStatLevel) test(difficulty int) (int, bool) {
	roll := this.roll()
	hits, botch := Dice.RollTest(roll, difficulty)
	if botch {
		this.awardXP(1)
	} else if hits <= this.Apt {
		this.awardXP(1)
	}
	return hits, botch
}

type Statted interface {
	GetStat(int) (int, int)
	GetStats([]int) (int, int)
	SetStatMod(int, string, int)
	ClearStatMod(int, string)
	ResetStatMods(int)
	AwardStatXP(int, int)
	RollStat(int) Roll
	TestStat(int, int) (int, bool)
}

type BasicStatted struct {
	Agility    StatLevel
	Endurance  StatLevel
	Health     StatLevel
	Reaction   StatLevel
	Strength   StatLevel
	Charisma   StatLevel
	Intuition  StatLevel
	Logic      StatLevel
	Perception StatLevel
	Willpower  StatLevel
}

func (this BasicStatted) dispatch(which int) StatLevel {
	switch which {
	case AGL:
		return this.Agility
	case END:
		return this.Endurance
	case HLT:
		return this.Health
	case REA:
		return this.Reaction
	case STR:
		return this.Strength
	case CHA:
		return this.Charisma
	case INT:
		return this.Intuition
	case LOG:
		return this.Logic
	case PER:
		return this.Perception
	case WP:
		return this.Willpower
	default:
		return nil
	}
}

func (this BasicStatted) GetStat(which int) (int, int) {
	return this.dispatch(which).get()
}

func (this BasicStatted) GetStats(which []int) (int, int) {
	value := 0
	shade := -1
	for _, i := range which {
		statValue, statShade := this.GetStat(i)
		value += statValue
		if statShade > shade {
			shade = statShade
		}
	}
	value = value / len(which)
	return value, shade
}

func (this BasicStatted) GetStatShade(which int) int { return this.dispatch(which).getShade() }

func (this *BasicStatted) SetStatMod(which int, mod string, value int) {
	this.dispatch(which).setMod(mod, value)
}

func (this *BasicStatted) ClearStatMod(which int, mod string) {
	this.dispatch(which).clearMod(mod)
}

func (this *BasicStatted) ResetStatMods(which int) {
	this.dispatch(which).resetMods()
}

func (this *BasicStatted) AwardStatXP(which int, award int) {
	this.dispatch(which).awardXP(award)
}

func (this BasicStatted) RollStat(which int) Roll {
	return this.dispatch(which).roll()
}

func (this BasicStatted) TestStat(which int, difficulty int) (int, bool) {
	return this.dispatch(which).test(difficulty)
}

func MakeStatLevel() StatLevel {
	result := BasicStatLevel{4, -1, make(map[string]int), 0, 7, 0}
	return &result
}

func MakeStatted() Statted {
	result := BasicStatted{
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
		MakeStatLevel(),
	}
	return &result
}
