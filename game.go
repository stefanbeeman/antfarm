package af

import (
  "strconv"
  "time"
)

type Game interface {
  tic()
  Start()
  Sleep(int)
  Stop()
  RunFor(int)
  addActor(Actor)
}

type BasicGame struct {
  World WorldState
  Now       int
  Materials map[string]Material
  Skills    map[string]Skill
  Actors    []Actor
  pacemaker *time.Ticker
}

func (this *BasicGame) tic() {
  this.Now++
  for _, actor := range this.Actors {
    // actor.tic(this)
  }
}

func (this *BasicGame) Start() {
  for t := range this.pacemaker.C {
    _ = t
    this.tic()
  }
}

func (this *BasicGame) Sleep(tics int) {
  duration, _ := time.ParseDuration(strconv.Itoa(tics*100) + "ms")
  time.Sleep(duration)
}

func (this *BasicGame) Stop() {
  this.pacemaker.Stop()
}

func (this *BasicGame) RunFor(tics int) {
  for i := 0; i < tics; i++ {
    this.tic()
  }
}

func (this *BasicGame) addActor(a Actor) {
  this.Actors = append(this.Actors, a)
}

func MakeGame(data string, width int, height int, worms int) Game {
  yml.setRoot(data)

  world := makeWorld(width, height)

  mats := yml.loadMaterials()
  skills := yml.loadSkills()
  units := make([]Actor, 0)
  pm := time.NewTicker(time.Millisecond * 100)
  Game := BasicGame{world, 0, mats, skills, units, pm}

  // for y, row := range Game. {
  //   for x, _ := range row {
  //     p := Point{x, y}
  //     c := makeCell(p, Game.Materials["rock"], false)
  //     if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
  //       c.setSolid(true)
  //     }
  //     Game.set(p, c)
  //   }
  // }

  return &Game
}