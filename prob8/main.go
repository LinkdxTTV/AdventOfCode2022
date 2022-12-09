package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	height  int
	visible bool
}

type Grid [][]Tree

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	grid := Grid{}

	for _, line := range inputSplit {
		if line == "" {
			continue
		}
		newLine := []Tree{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(line, "failed to convert num")
			}
			newLine = append(newLine, Tree{num, false})
		}
		grid = append(grid, newLine)
	}

	width := len(grid[0]) // Horizontal
	length := len(grid)   // Vertical

	visible := 0

	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if i == 0 || j == 0 || i == width-1 || j == length-1 {
				visible++
				continue
			}
			if grid.calculateVisibility(i, j) {
				visible++
			}
		}
	}

	fmt.Println(visible)

	highestScenicSoFar := 0
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if i == 0 || j == 0 || i == width-1 || j == length-1 {
				continue // Just assume it isnt on the edge
			}

			scenicScore := grid.calculateScenicScore(i, j)
			if scenicScore > highestScenicSoFar {
				highestScenicSoFar = scenicScore
			}
		}
	}

	fmt.Println(highestScenicSoFar)
}

func (g Grid) calculateVisibility(x, y int) bool {
	width := len(g[0])
	length := len(g)

	tree := g[y][x]
	// Left
	for i := x - 1; i >= 0; i-- {
		if g[y][i].height >= tree.height {
			break
		}
		if i == 0 {
			tree.visible = true
			return true
		}
	}

	// Right
	for i := x + 1; i < width; i++ {
		if g[y][i].height >= tree.height {
			break
		}
		if i == width-1 {
			tree.visible = true
			return true
		}
	}

	// Up
	for j := y - 1; j >= 0; j-- {
		if g[j][x].height >= tree.height {
			break
		}
		if j == 0 {
			tree.visible = true
			return true
		}
	}

	// Down
	for j := y + 1; j < length; j++ {
		if g[j][x].height >= tree.height {
			break
		}
		if j == length-1 {
			tree.visible = true
			return true
		}
	}
	return false
}

func (g Grid) calculateScenicScore(x, y int) int {
	width := len(g[0])
	length := len(g)

	tree := g[y][x]
	// Left
	counterLeft := 1
	for i := x - 1; i > 0; i-- {
		if g[y][i].height >= tree.height {
			break
		}
		counterLeft++
	}

	// Right
	counterRight := 1
	for i := x + 1; i < width-1; i++ {
		if g[y][i].height >= tree.height {
			break
		}
		counterRight++
	}

	// Up
	counterUp := 1
	for j := y - 1; j > 0; j-- {
		if g[j][x].height >= tree.height {
			break
		}
		counterUp++
	}

	// Down
	counterDown := 1
	for j := y + 1; j < length-1; j++ {
		if g[j][x].height >= tree.height {
			break
		}
		counterDown++
	}

	// fmt.Println(x, y, tree.height, "LRUD:", counterLeft, counterRight, counterUp, counterDown)
	return counterUp * counterDown * counterLeft * counterRight

}
