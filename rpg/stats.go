package rpg

type Stat interface {
	get() (int, int)
	getBase() int
	setBase(int)
	getShade() int
	setShade(int)
	getMod(string) int
	setMod(string, int)
	clearMod(string)
	resetMods(string)
	getMax() int
	setMax(int)
	getXP() int
	awardXP(int)
}

type BasicStat struct {
	Base  int
	Shade int
	Mods  map[string]int
	Max   int
	XP    int
}

func (this BasicStat) get() (int, int) {
	value := this.Base
	for _, mod := range this.Mods {
		value += mod
	}
	return value, this.Shade
}

func (this BasicStat) getBase() int {
	return this.Base
}

func (this *BasicStat) setBase(value int) {
	this.Base = value
}

func (this BasicStat) getShade() int {
	return this.Shade
}

func (this *BasicStat) setShade(value int) {
	this.Shade = value
}

func (this BasicStat) getMod(mod string) int {
	return this.Mods[mod]
}

func (this *BasicStat) setMod(mod string, value int) {
	this.Mods[mod] = value
}

func (this *BasicStat) clearMod(mod string) {
	this.Mods[mod] = 0
}

func (this *BasicStat) resetMods(mod string, value int) {
	this.Mods = make(map[string]int)
}

func (this BasicStat) getMax() int {
	return this.Max
}

func (this *BasicStat) setMax(value int) {
	this.Max = value
}

func (this BasicStat) getXP() int {
	return this.XP
}

func (this *BasicStat) awardXP(value int) {
	this.XP += value
	if this.Base < this.Max {
		next := this.Base + 1
		if this.XP >= (next * STATXP) {
			this.Base += 1
			this.XP = 0
		}
	}
}

type Statted interface {
	GetStat(int) (int, int)
	GetStats([]int) (int, int)
	SetStatMod(int, string, int)
	ClearStatMod(int, string)
	ResetStatMods(int)
	AwardStatXP(int, int)
	TestStat(int) int
	TestStats([]int) int
}

type BasicStatted struct {
	Agility    Stat
	Endurance  Stat
	Health     Stat
	Reaction   Stat
	Strength   Stat
	Charisma   Stat
	Intuition  Stat
	Logic      Stat
	Perception Stat
	Willpower  Stat
}

func (this BasicStatted) dispatch(which int) Stat {
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
		statValue, statShade = this.GetStat(i)
		value += statValue
		if statShade > shade {
			shade = statShade
		}
	}
	value = value / len(which)
	return value, shade
}

func (this BasicStatted) GetStatShade(which int) int {
	return this.dispatch(which).getShade()
}

func (this *BasicStatted) SetStatMod(which int, mod string, value int) {
	this.dispatch(which).setMod(mod, value)
}

func (this *BasicStatted) ClearStatMod(which int, mod string) {
	this.dispatch(which).clearMod(mod)
}

func (this *BasicStatted) ResetStatMods(which int, mod string) {
	this.dispatch(which).resetMods(mod)
}

func (this *BasicStatted) AwardStatXP(which int, award int) {
	this.dispatch(which).awardXP(award)
}

func (this *BasicStatted) TestStat(which int) {
	dice, shade := this.GetStat(which)
	return Dice.RollDice(dice, 6, shade)
}

func (this *BasicStatted) TestStats(which int) {
	dice, shade := this.GetStats(which)
	return Dice.RollDice(dice, 6, shade)
}

func MakeStat() Stat {
	result := BasicStat{4, -1, make(map[string]int), 10, 0}
	return &result
}

func MakeStatted() Stat {
	result := BasicStatted{
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
		MakeStat(),
	}
}
