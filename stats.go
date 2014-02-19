package af

var STAT_ADJ = [12]string{"pitiable", "poor", "mediocre", "average", "good", "great", "superb", "heroic", "fantastic", "epic", "legendary", "mythic"}
var STAT_XP = [12]int{2, 10, 25, 45, 70, 100, 135, 175, 220, 270, 325, 385}

type Stat interface {
	get() int
	getBase() int
	setBase(int)
	getMod() int
	setMod(int)
	getMax() int
	setMax(int)
	getXP() int
	setXP(int)
	awardXP(int)
	adj() string
}

type BasicStat struct {
	base int
	mod  int
	max  int
	xp   int
}

func (this BasicStat) get() int {
	return this.base + this.mod
}

func (this BasicStat) getBase() int {
	return this.base
}

func (this *BasicStat) setBase(value int) {
	this.base = value
}

func (this BasicStat) getMod() int {
	return this.mod
}

func (this *BasicStat) setMod(value int) {
	this.mod = value
}

func (this BasicStat) getMax() int {
	return this.mod
}

func (this *BasicStat) setMax(value int) {
	this.mod = value
}

func (this BasicStat) getXP() int {
	return this.xp
}

func (this BasicStat) setXP(xp int) {
	this.xp = xp
	this.calculateXP()
}

func (this *BasicStat) awardXP(award int) {
	this.xp += award
	this.calculateXP()
}

func (this BasicStat) calculateXP() {
	value := this.getBase()
	if value < len(STAT_XP) && value >= 0 && value < this.max {
		next := STAT_XP[value]
		if this.xp >= next {
			this.base += 1
			this.calculateXP()
		}
	}
}

func (this BasicStat) adj() string {
	value := this.get()
	if value >= len(STAT_ADJ) {
		return "inhuman"
	} else if value < 0 {
		return "nonexistant"
	} else {
		return STAT_ADJ[this.get()]
	}
}

type PhysicalStatline struct {
	Body     Stat
	Agility  Stat
	Reaction Stat
	Strength Stat
}

func (this PhysicalStatline) physicalLimit() int {
	bod := this.Body.get()
	rea := this.Reaction.get()
	str := this.Strength.get()
	return (bod + rea + (str * 2)) / 3
}

type MentalStatline struct {
	Willpower Stat
	Logic     Stat
	Intuition Stat
	Charisma  Stat
}

func (this MentalStatline) mentalLimit() int {
	wp := this.Willpower.get()
	log := this.Logic.get()
	intu := this.Intuition.get()
	return (wp + intu + (log * 2)) / 3
}

type SpiritualStatline struct {
	Peity      Stat
	Magic      Stat
	Corruption Stat
	Insanity   Stat
}

func (this SpiritualStatline) magicalLimit() int {
	mag := this.Magic.get()
	cor := this.Corruption.get()
	ins := this.Insanity.get()
	return ((mag * 2) + ins + cor) / 3
}

type Statline struct {
	PhysicalStatline
	MentalStatline
	SpiritualStatline
}
