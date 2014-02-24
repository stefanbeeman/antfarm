package af

type stat int

// The ten basic stats.
const (
	ST = 0 // Strength, the raw force the unit can exert on the world.
	AG = 1 // Agility, the precision and grace with which the unit moves.
	TO = 2 // Toughness, how resistant the unit is to damage.
	EN = 3 // Endurance, how quickly the unit tires.
	HT = 4 // Health, how quickly the unit recovers from injury, poison, and disease.
	RE = 6 // Reaction, how quickly the unit responds to new circumstances.
	WP = 5 // Willpower, how determined and driven the unit is.
	LG = 7 // Logic, the unit's ability to think abstractly.
	IN = 8 // Intuiton, the unit's awarnesss and sensitivity.
	CH = 9 // Charisma, the unit's presence and social ability.
)

type Statline interface {
	get(stat) int
	set(stat, int)
	addTo(stat, int)
}

type BasicStatline struct {
	Strength  int
	Agility   int
	Toughness int
	Endurance int
	Health    int
	Reaction  int
	Willpower int
	Logic     int
	Intuition int
	Charisma  int
}

func (this BasicStatline) get(which stat) int {
	switch which {
	case ST:
		return this.Strength
	case AG:
		return this.Agility
	case TO:
		return this.Toughness
	case EN:
		return this.Endurance
	case HT:
		return this.Health
	case RE:
		return this.Reaction
	case WP:
		return this.Willpower
	case LG:
		return this.Logic
	case IN:
		return this.Intuition
	case CH:
		return this.Charisma
	default:
		return 0
	}
}

func (this *BasicStatline) set(which stat, value int) {
	switch which {
	case ST:
		this.Strength = value
	case AG:
		this.Agility = value
	case TO:
		this.Toughness = value
	case EN:
		this.Endurance = value
	case HT:
		this.Health = value
	case RE:
		this.Reaction = value
	case WP:
		this.Willpower = value
	case LG:
		this.Logic = value
	case IN:
		this.Intuition = value
	case CH:
		this.Charisma = value
	}
}

func (this *BasicStatline) addTo(which stat, value int) {
	switch which {
	case ST:
		this.Strength += value
	case AG:
		this.Agility += value
	case TO:
		this.Toughness += value
	case EN:
		this.Endurance += value
	case HT:
		this.Health += value
	case RE:
		this.Reaction += value
	case WP:
		this.Willpower += value
	case LG:
		this.Logic += value
	case IN:
		this.Intuition += value
	case CH:
		this.Charisma += value
	}
}

type att int

// The five combat attributes
const (
	ACU = 0 // Acuity (AG + RE)/2, a measure of the unit's ability to meaningfully react in melee combat.
	AIM = 1 // Aim (AG + IN)/2, a measure of the unit's spatial awareness important to ranged combat.
	KDN = 2 // Knockdown (ST + AG)/2, how much oomph is required to knock this unit prone.
	KOT = 3 // Knockout (TO + WP/2), how much pain the unit can take before passing out.
	MOV = 4 // Movement (ST + AG + EN)/2, how quickly the unit moves, relative to others of the same unit type.
)

// The five magic attributes
const (
	RAU = 5 // Rauch (WP + IN + CH)/2, how much magical power the unit can direct at once
	FRM = 6 // Form (RE + IN)/2, the unit's control over magical energy.
	ART = 7 // Art (EN + LG)/2, the unit's ability to safely endure the side-effects of magic.
	DIS = 8 // Discipline (EN + WP)/2, the unit's caution and precision in peforming rituals.
	DRW = 9 // Draw (HT + WP)/2, the rate at which the unit can draw magical energy into itself.
)

type Statted interface {
	getStat(stat) int
	getAtt(att) int
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
}

type BasicStatted struct {
	BaseStats Statline
	StatMods  Statline
	StatXP    Statline
}

func (this BasicStatted) getStat(which stat) int {
	base := this.BaseStats.get(which)
	mod := this.StatMods.get(which)
	return (base + mod)
}

func (this BasicStatted) getAtt(which att) int {
	switch which {
	case ACU:
		return (this.getStat(AG) + this.getStat(RE)) / 2
	case AIM:
		return (this.getStat(AG) + this.getStat(IN)) / 2
	case KDN:
		return (this.getStat(ST) + this.getStat(AG)) / 2
	case KOT:
		return (this.getStat(TO) + this.getStat(WP)/2)
	case MOV:
		return (this.getStat(ST) + this.getStat(AG) + this.getStat(EN)) / 2
	case RAU:
		return (this.getStat(WP) + this.getStat(IN) + this.getStat(CH)) / 2
	case FRM:
		return (this.getStat(RE) + this.getStat(IN)) / 2
	case ART:
		return (this.getStat(EN) + this.getStat(LG)) / 2
	case DIS:
		return (this.getStat(EN) + this.getStat(WP)) / 2
	case DRW:
		return (this.getStat(HT) + this.getStat(WP)) / 2
	default:
		return 0
	}
}

func (this BasicStatted) getBaseStat(which stat) int {
	return this.BaseStats.get(which)
}

func (this *BasicStatted) setBaseStat(which stat, value int) {
	this.BaseStats.set(which, value)
}

func (this BasicStatted) getStatMod(which stat) int {
	return this.StatMods.get(which)
}

func (this *BasicStatted) setStatMod(which stat, value int) {
	this.StatMods.set(which, value)
}

func (this *BasicStatted) addStatMod(which stat, value int) {
	this.StatMods.addTo(which, value)
}

func (this *BasicStatted) resetStatMod(which stat) {
	this.StatMods.set(which, 0)
}

func (this *BasicStatted) resetStatMods() {
	this.StatMods = new(BasicStatline)
}

func (this BasicStatted) getStatXP(which stat) int {
	return this.StatXP.get(which)
}

func (this *BasicStatted) setStatXP(which stat, value int) {
	this.StatXP.set(which, value)
	this.calculateStatXP(which)
}

func (this *BasicStatted) addStatXP(which stat, value int) {
	this.StatMods.addTo(which, value)
	this.calculateStatXP(which)
}

func (this *BasicStatted) calculateStatXP(which stat) {
	current := this.getBaseStat(which)
	next := current + 1
	nextXP := next * 5
	if this.getStatXP(which) > nextXP {
		this.setBaseStat(which, next)
		this.StatXP.set(which, this.getStatXP(which)-nextXP)
	}
}
