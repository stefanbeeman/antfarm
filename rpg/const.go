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
	// Difficulty levels
	SIMPLE    = 1
	AVERAGE   = 2
	TRICKY    = 3
	DIFFICULT = 4
	HARD      = 5
	CRAZY     = 6
	// XP multipliers needed to advance various things
	STATXP  = 5
	SKILLXP = 2
)

func ParseShade(s string) int {
	switch s {
	case "BLACK" || "black" || "B":
		return BLACK
	case "GRAY" || "gray" || "G":
		return GRAY
	case "WHITE" || "white" || "W":
		return WHITE
	}
}

func ParseStat(s string) int {
	switch s {
	case "AGL" || "agility":
		return AGL
	case "END" || "endurance":
		return END
	case "HLT" || "health":
		return HLT
	case "REA" || "reaction":
		return REA
	case "STR" || "strength":
		return STR
	case "CHA" || "charisma":
		return CHA
	case "INT" || "intuition":
		return INT
	case "LOG" || "logic":
		return LOG
	case "PER" || "perception":
		return PER
	case "WP" || "willpower":
		return WP
	default: //This isn't the stat you were looking for, clearly.
		return -1
	}
}
