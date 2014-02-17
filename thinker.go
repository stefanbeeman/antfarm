package antfarm

import (
  "math/rand"
  "fmt"
)

type Thinker interface {
	think(*BasicUnit, *World) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *BasicUnit, w *World) Task {
  task := NO_TASK
  for task.complete() {
    d := Point{rand.Intn(20),rand.Intn(20)}
    fmt.Println("I'm heading to ", d)
    task = BasicTask{
      func() bool {
        return u.Position.equals(d)
      },
      u.mover.moveTo(u, d),
    }
  }

  return task
}