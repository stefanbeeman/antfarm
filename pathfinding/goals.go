package pathfinding

import . "github.com/stefanbeeman/antfarm/common"

type BasicGoal struct {
	Location
	WeightValue int
}

func (this BasicGoal) Weight() int { return this.WeightValue }

type BasicGoalDecider struct {
	goals []Goal
}

func (this *BasicGoalDecider) AddGoals(goals []Goal) {
	this.goals = append(this.goals, goals...)
}

func (this *BasicGoalDecider) RemoveGoals(goals []Goal) {
	removed := make(map[Goal]bool)
	for _, g := range goals {
		removed[g] = true
	}
	modified := []Goal{}
	for _, g := range this.goals {
		if _, ok := removed[g]; ok {
			modified = append(modified, g)
		}
	}
	this.goals = modified
}

func (this *BasicGoalDecider) BestGoal() Goal {
	return this.goals[0]
}

func (this *BasicGoalDecider) H(p Location) int {
	return 10 * p.DistanceTo(this.BestGoal())
}

func MakeGoalDecider() GoalDecider { return &BasicGoalDecider{[]Goal{}} }
