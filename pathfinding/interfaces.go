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
}

type MovementAlg interface {
	RegisterMoveCost(func(Location) (int, bool))
	NextStep(Location, Location) (Location, bool)
	findPath(Location, Location) bool
}
