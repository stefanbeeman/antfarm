package antfarm

import (
	"bufio"
	"fmt"
	"github.com/stefanbeeman/antfarm/world"
	"os"
	"strconv"
	"strings"
)

func (this BasicGame) StartShell() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Antfarm!")
	for {
		fmt.Print("AF: ")
		raw, _ := reader.ReadString('\n')
		raw = strings.TrimSuffix(raw, "\n")
		commands := strings.Split(raw, " ")
		if commands[0] == "exit" {
			fmt.Println("later")
			break
		} else {
			this.RunCommand(commands)
		}
	}
}

func (this BasicGame) RunCommand(commands []string) {
	switch commands[0] {
	case "runfor":
		tics, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println("try a real number")
		} else {
			this.RunFor(tics)
		}
	case "now":
		fmt.Println(this.Now)
	case "save":
		if len(commands) > 1 {
			this.Save(commands[1])
		} else {
			fmt.Println("Please provide a filename for your save.")
		}
	case "maze":
		w, _ := strconv.Atoi(commands[1])
		h, _ := strconv.Atoi(commands[2])
		t, _ := strconv.Atoi(commands[3])
		s, _ := strconv.Atoi(commands[4])
		maze := world.GenMaze(w, h, t, s)
		world.PrintMaze(maze)
	default:
		fmt.Println("try that again")
	}
}
