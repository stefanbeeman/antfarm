package pathfinding

import . "github.com/stefanbeeman/antfarm/common"

type Goal interface {
	Location
	Weight() int
}

type GoalDecider interface {
	AddGoals([]Goal)
	RemoveGoals([]Goal)
	BestGoal() Goal
	H(Location) int
}

type MovementAlg interface {
	GoalDecider
	NextStep(Location) (Location, bool)
	RegisterMoveCost(func(Location) (int, bool))
	findPath(Point) bool
}
