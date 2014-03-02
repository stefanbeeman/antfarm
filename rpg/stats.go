package rpg

// The shades
const (
	BLACK = -1
	GRAY  = 0
	WHITE = 1
)

// The stats
const (
	AGL = 0 // Agility
	END = 1 // Endurance
	HLT = 2 // Health
	REA = 3 // Reaction
	STR = 4 // Strength
	CHA = 5 // Charisma
	INT = 6 // Intuition
	LOG = 7 // Logic
	PER = 8 // Perception
	WIL = 9 // Willpower
)

// XP multiplier needed to advance a stat
const (
	STATXP = 5
)

func parseStat(s string) int {
	switch s {
	case "AGL":
		return 0
	case "END":
		return 1
	case "HLT":
		return 2
	case "REA":
		return 3
	case "STR":
		return 4
	case "CHA":
		return 5
	case "INT":
		return 6
	case "LOG":
		return 7
	case "PER":
		return 8
	case "WIL":
		return 9
	default:
		return 0
	}
}

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
