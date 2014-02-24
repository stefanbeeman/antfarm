package af

import (
	"strconv"
	"time"
)

type World interface {
	tic()
	Start()
	Sleep(int)
	Stop()
	RunFor(int)
	addActor(Actor)
}

type BasicWorld struct {
	Grid2D
	Now       int
	Materials map[string]Material
	Skills    map[string]Skill
	Actors    []Actor
	pacemaker *time.Ticker
}

func (this *BasicWorld) tic() {
	this.Now++
	for _, actor := range this.Actors {
		actor.tic(this)
	}
}

func (this *BasicWorld) Start() {
	for t := range this.pacemaker.C {
		_ = t
		this.tic()
	}
}

func (this *BasicWorld) Sleep(tics int) {
	duration, _ := time.ParseDuration(strconv.Itoa(tics*100) + "ms")
	time.Sleep(duration)
}

func (this *BasicWorld) Stop() {
	this.pacemaker.Stop()
}

func (this *BasicWorld) RunFor(tics int) {
	for i := 0; i < tics; i++ {
		this.tic()
	}
}

func (this *BasicWorld) addActor(a Actor) {
	this.Actors = append(this.Actors, a)
}

func MakeWorld(data string, width int, height int, worms int) World {
	yml.setRoot(data)
	grid := makeGrid2D(width, height)
	mats := yml.loadMaterials()
	skills := yml.loadSkills()
	units := make([]Actor, 0)
	pm := time.NewTicker(time.Millisecond * 100)
	world := BasicWorld{grid, 0, mats, skills, units, pm}

	for y, row := range world.Cells {
		for x, _ := range row {
			p := Point{x, y}
			c := makeCell(p, world.Materials["rock"], false)
			if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
				c.setSolid(true)
			}
			world.set(p, c)
		}
	}

	for n := 0; n < worms; n++ {
		rp := world.random()
		a := makeWorm(rp)
		world.addActor(a)
	}
	return &world
}
