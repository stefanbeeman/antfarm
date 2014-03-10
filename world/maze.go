package world

import (
	"fmt"
	. "github.com/stefanbeeman/antfarm/common"
	"math/rand"
)

func initMaze(width int, height int) [][]bool {
	maze := make([][]bool, height)
	for y := 0; y < height; y++ {
		maze[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			maze[y][x] = true
		}
	}
	return maze
}

func contains(maze [][]bool, x int, y int) bool {
	if y >= 1 && y < len(maze)-1 {
		if x >= 1 && x < len(maze[0])-1 {
			return true
		}
	}
	return false
}

func canCut(maze [][]bool, x int, y int, thresh int, search int) bool {
	if contains(maze, x, y) {
		sum := 0
		for i := x - search; i <= x+search; i++ {
			for j := y - search; j <= y+search; j++ {
				if !contains(maze, i, j) || maze[j][i] {
					sum++
				}
			}
		}
		return sum > thresh
	}
	return false
}

func visit(visited map[Location]bool, stack []Location, neigh Location) []Location {
	if _, ok := visited[neigh]; !ok {
		visited[neigh] = true
		stack = append(stack, neigh)
	}
	return stack
}

func GenMaze(width int, height int, thresh int, search int) [][]bool {
	maze := initMaze(width, height)
	start := Point{1, 1}
	stack := make([]Location, 0)
	visited := make(map[Location]bool)
	visited[start] = true
	stack = append(stack, start)
	for len(stack) > 0 {
		i := len(stack) - 1
		current := stack[i]
		x, y := current.Coords()
		stack = stack[:i]
		if canCut(maze, x, y, thresh, search) {
			maze[y][x] = false
			toVisit := current.Neighbors()
			for len(toVisit) > 0 {
				blah := rand.Intn(len(toVisit))
				stack = visit(visited, stack, toVisit[blah])
				toVisit = append(toVisit[:blah], toVisit[blah+1:]...)
			}
		}
	}
	return maze
}

func PrintMaze(maze [][]bool) {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {
			if maze[y][x] {
				fmt.Print("[]")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}
