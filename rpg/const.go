package rpg

const (
	// The shades
	BLACK = -1 // Black or "mundane" shade.
	GRAY  = 0  // Gray or "heroic" shade.
	WHITE = 1  // White or "supernatural" shade.
	// The stats
	AGL = 0 // Agility
	END = 1 // Endurance
	HLT = 2 // Health
	REA = 3 // Reaction
	STR = 4 // Strength
	CHA = 5 // Charisma
	INT = 6 // Intuition
	LOG = 7 // Logic
	PER = 8 // Perception
	WP  = 9 // Willpower
	// Skill levels
	UNTRAINED   = 0
	BEGINNER    = 1
	NOVICE      = 2
	APPRENTICE  = 3
	COMPETENT   = 4
	JOURNEYMAN  = 5
	MASTER      = 6
	GRANDMASTER = 7
	LEGENDARY   = 8
	FLAWLESS    = 9
	// Difficulty levels
	SIMPLE    = 1
	AVERAGE   = 2
	TRICKY    = 3
	DIFFICULT = 4
	HARD      = 5
	INSANE    = 6
	// XP multipliers needed to advance various things
	STATXP  = 5
	SKILLXP = 2
	// Wound levels
	GRAZE   = 0
	MINOR   = 1
	SERIOUS = 2
	MAJOR   = 3
	SEVERE  = 4
	MORTAL  = 5
)

func ParseShade(s string) int {
	switch s {
	case "GRAY":
		return GRAY
	case "WHITE":
		return WHITE
	default:
		return BLACK
	}
}

func ParseStat(s string) int {
	switch s {
	case "AGL":
		return AGL
	case "END":
		return END
	case "HLT":
		return HLT
	case "REA":
		return REA
	case "STR":
		return STR
	case "CHA":
		return CHA
	case "INT":
		return INT
	case "LOG":
		return LOG
	case "PER":
		return PER
	case "WP":
		return WP
	default: //This isn't the stat you were looking for, clearly.
		return -1
	}
}
