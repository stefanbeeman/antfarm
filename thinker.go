package af

import (
	"fmt"
	"math/rand"
)

const (
  OBJECTIVES int = 2
  MAX_WEIGHT int = 1
)

type Thinker interface {
	think(*BasicUnit, World) Task
}

type BasicThinker struct{}

func (this BasicThinker) think(u *BasicUnit, w World) Task {
	task := NO_TASK
	for task.complete() {
    desires := []DesirePoint{}
    for i := 0; i < OBJECTIVES; i++ {
      desires = append(desires, DesirePoint{w.random(), float64(rand.Intn(MAX_WEIGHT) + 1)})
    }
		fmt.Println("I'm heading to ", desires)
		task = BasicTask{
			func() bool {
        for _, d := range desires {
          if u.Position.equals(d.Position) { return true }
        }
				return false
			},
			u.mover.moveTo(u, desires),
		}
	}

	return task
}
