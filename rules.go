package antfarm

//import (
//	"fmt"
//	"math/rand"
//)

type statline struct {
	Body      int
	Agility   int
	Reaction  int
	Strength  int
	Willpower int
	Logic     int
	Intuition int
	Charisma  int
	Magic     int
	Essence   int
}

type Skill struct {
	Stat     string
	Limit    string
	Type     string
	Default  bool
	Abstract bool
	Title    string
}
