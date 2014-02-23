package af

// The base eight stats.
type stat int

const (
	BOD = 0
	AGL = 1
	REA = 2
	STR = 3
	WP  = 4
	LOG = 5
	INT = 6
	CHA = 7
)

// The three limits
type limit int

const (
	PHYSICAL = 0
	MENTAL   = 1
	SOCIAL   = 2
)

type Statline interface {
	get(stat) int
	set(stat, int)
	addTo(stat, int)
}

type BasicStatline struct {
	Body      int
	Agility   int
	Reaction  int
	Strength  int
	Willpower int
	Logic     int
	Intuition int
	Charisma  int
}

func (this BasicStatline) get(which stat) int {
	switch which {
	case BOD:
		return this.Body
	case AGL:
		return this.Agility
	case REA:
		return this.Reaction
	case STR:
		return this.Strength
	case WP:
		return this.Willpower
	case LOG:
		return this.Logic
	case INT:
		return this.Intuition
	case CHA:
		return this.Charisma
	default:
		return 0
	}
}

func (this *BasicStatline) set(which stat, value int) {
	switch which {
	case BOD:
		this.Body = value
	case AGL:
		this.Agility = value
	case REA:
		this.Reaction = value
	case STR:
		this.Strength = value
	case WP:
		this.Willpower = value
	case LOG:
		this.Logic = value
	case INT:
		this.Intuition = value
	case CHA:
		this.Charisma = value
	}
}

func (this *BasicStatline) addTo(which stat, value int) {
	switch which {
	case BOD:
		this.Body += value
	case AGL:
		this.Agility += value
	case REA:
		this.Reaction += value
	case STR:
		this.Strength += value
	case WP:
		this.Willpower += value
	case LOG:
		this.Logic += value
	case INT:
		this.Intuition += value
	case CHA:
		this.Charisma += value
	}
}

type Stats interface {
	getStat(stat) int
	getBaseStat(stat) int
	setBaseStat(stat, int)
	getStatMod(stat) int
	setStatMod(stat, int)
	addStatMod(stat, int)
	resetStatMod(stat)
	resetStatMods()
	getStatXP(stat) int
	setStatXP(stat, int)
	addStatXP(stat, int)
	getLimit(limit)
}

type BasicStats struct {
	BaseStats Statline
	StatMods  Statline
	StatXP    Statline
}

func (this BasicStats) getStat(which stat) int {
	base := this.BaseStats.get(which)
	mod := this.StatMods.get(which)
	return (base + mod)
}

func (this BasicStats) getBaseStat(which stat) int {
	return this.BaseStats.get(which)
}

func (this *BasicStats) setBaseStat(which stat, value int) {
	this.BaseStats.set(which, value)
}

func (this BasicStats) getStatMod(which stat) int {
	return this.StatMods.get(which)
}

func (this *BasicStats) setStatMod(which stat, value int) {
	this.StatMods.set(which, value)
}

func (this *BasicStats) addStatMod(which stat, value int) {
	this.StatMods.addTo(which, value)
}

func (this *BasicStats) resetStatMod(which stat) {
	this.StatMods.set(which, 0)
}

func (this *BasicStats) resetStatMods() {
	this.resetStatMod(BOD)
	this.resetStatMod(AGL)
	this.resetStatMod(REA)
	this.resetStatMod(STR)
	this.resetStatMod(WP)
	this.resetStatMod(LOG)
	this.resetStatMod(INT)
	this.resetStatMod(CHA)
}

func (this BasicStats) getStatXP(which stat) int {
	return this.StatXP.get(which)
}

func (this *BasicStats) setStatXP(which stat, value int) {
	this.StatXP.set(which, value)
	this.calculateStatXP(which)
}

func (this *BasicStats) addStatXP(which stat, value int) {
	this.StatMods.addTo(which, value)
	this.calculateStatXP(which)
}

func (this *BasicStats) calculateStatXP(which stat) {
	current := this.getBaseStat(which)
	next := current + 1
	nextXP := next * 5
	if this.getStatXP(which) > nextXP {
		this.setBaseStat(which, next)
		this.StatXP.set(which, this.getStatXP(which)-nextXP)
	}
}

func (this BasicStats) getLimit(which limit) int {
	switch which {
	case PHYSICAL:
		bod := this.getStat(BOD)
		rea := this.getStat(REA)
		str := this.getStat(STR)
		return (bod + rea + (str * 2)) / 3
	case MENTAL:
		wp := this.getStat(WP)
		log := this.getStat(LOG)
		intu := this.getStat(INT)
		return (wp + intu + (log * 2)) / 3
	case SOCIAL:
		wp := this.getStat(WP)
		intu := this.getStat(INT)
		cha := this.getStat(CHA)
		return (wp + intu + (cha * 2)) / 3
	default:
		return 0
	}
}
