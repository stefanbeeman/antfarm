package af

import (
	"fmt"
)

type Action interface {}

type BasicAction struct {}

type Goal interface {
  Location
  Weight() int
}

type BasicGoal struct {
  Location
  weight int
}

func (this BasicGoal) Weight() { return this.weight }

type Thinker interface {
  Init(Unit)
  Think(Unit)
}

type BasicThinker struct{
  goals []Goal
}

func (this BasicThinker) Init(u Unit) {
  u.AddGoals( this.goals )
}

func (this BasicThinker) Think(u Unit) {
  u.SetAction( u.Move(u) )
}