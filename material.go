package antfarm

import (
	"fmt"
	"github.com/soundcloud/goyaml"
	"io/ioutil"
)

const ROCK = 0
const DIRT = 1
const SAND = 2
const STONE = 3
const SANDSTONE = 4
const MARBLE = 5
const OBSIDIAN = 6
const ORECOPPER = 7
const OREIRON = 8
const ORESILVER = 9
const OREGOLD = 10
const OREORICHALCUM = 11
const ICE = 12
const FLESH = 13
const CRYSTAL = 14
const BONE = 15

type Material struct {
	Name      string
	Rune      rune
	Structure int
	Armor     int
}

func LoadMaterials() []Material {
	mats := make([]Material, 0)
	buffer, err := ioutil.ReadFile("app/data/material.yml")
	if err != nil {
		fmt.Println(err)
	} else {
		err = goyaml.Unmarshal(buffer, &mats)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(mats)
	}
	return mats
}
