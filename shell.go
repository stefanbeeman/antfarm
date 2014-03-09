package antfarm

import (
	"bufio"
	"fmt"
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
	default:
		fmt.Println("try that again")
	}
}
