package af

type Statline interface {
	getStat(string) int
	getBaseStat(string) int
	setBaseStat(string, int)
	getModStat(string) int
	setModStat(string, int)
	getMaxStat(string) int
	setMaxStat(string, int)
	getStatXP(string) int
	setStatXP(string, int)
	awardStatXP(string, int)
	getLimit(string) int
	statTest(string, string, string) (int, bool)
	statTitle(string) string
	addStat(string, int, int, string)
}

type BasicStatline struct {
	Base map[string]int
	Mod  map[string]int
	Max  map[string]int
	XP   map[string]int
}

func (this BasicStatline) getStat(stat string) int {
	return (this.Base[stat] + this.Mod[stat])
}

func (this BasicStatline) getBaseStat(stat string) int {
	return this.Base[stat]
}

func (this *BasicStatline) setBaseStat(stat string, value int) {
	this.Base[stat] = value
}

func (this BasicStatline) getModStat(stat string) int {
	return this.Mod[stat]
}

func (this *BasicStatline) setModStat(stat string, value int) {
	this.Mod[stat] = value
}

func (this BasicStatline) getMaxStat(stat string) int {
	return this.Mod[stat]
}

func (this BasicStatline) setMaxStat(stat string, value int) {
	this.Max[stat] = value
}

func (this BasicStatline) getStatXP(stat string) int {
	return this.XP[stat]
}

func (this *BasicStatline) setStatXP(stat string, value int) {
	this.XP[stat] = value
	this.calculateStatXP(stat)
}

func (this *BasicStatline) awardStatXP(stat string, award int) {
	this.XP[stat] += award
	this.calculateStatXP(stat)
}

func xpForStat(rating int) int {
	xp := 0
	for i := rating; i > 0; i-- {
		xp += (i * 5)
	}
	return xp
}

func (this *BasicStatline) calculateStatXP(stat string) {
	base := this.Base[stat]
	max := this.Max[stat]
	xp := this.XP[stat]
	if base > 0 && base <= max && xp > xpForStat(base+1) {
		this.Base[stat] = this.Base[stat] + 1
	}
}

func (this BasicStatline) getLimit(limit string) int {
	switch limit {
	case "physical":
		bod := this.getStat("body")
		rea := this.getStat("reaction")
		str := this.getStat("strength")
		return (bod + rea + (str * 2)) / 3
	case "mental":
		wp := this.getStat("willpower")
		log := this.getStat("logic")
		intu := this.getStat("intuition")
		return (wp + (log * 2) + intu) / 3
	case "social":
		wp := this.getStat("willpower")
		intu := this.getStat("intuition")
		cha := this.getStat("charisma")
		return (wp + intu + (cha * 2)) / 3
	case "arcane":
		mag := this.getStat("magic")
		san := this.getStat("insanity")
		cor := this.getStat("corruption")
		return ((mag * 2) + san + cor) / 2
	default:
		return 0
	}
}

func (this BasicStatline) statTest(first string, second string, whichLimit string) (int, bool) {
	dice := this.getStat(first) + this.getStat(second)
	limit := this.getLimit(whichLimit)
	afd.rollDice(dice, limit)
}

func (this StatTitle) statTitle()
