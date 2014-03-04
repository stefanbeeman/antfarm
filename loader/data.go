package loader

import (
	"github.com/soundcloud/goyaml"
	. "github.com/stefanbeeman/antfarm/common"
	"io/ioutil"
)

func loadFile(path string, pointer interface{}) {
	Console.Meh("Loading " + path)
	buffer, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		Console.Broke(readErr)
	}
	marshErr := goyaml.Unmarshal(buffer, pointer)
	if marshErr != nil {
		Console.Broke(marshErr)
	}
}

func LoadData(root string) {
	loadMaterials(root)
	loadSkills(root)
}
