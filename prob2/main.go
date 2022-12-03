package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var shapePoints map[string]int = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

var outcomePoints map[string]int = map[string]int{
	"A X": 3,
	"B X": 0,
	"C X": 6,
	"A Y": 6,
	"B Y": 3,
	"C Y": 0,
	"A Z": 0,
	"B Z": 6,
	"C Z": 3,
}

var part2OutcomePoints map[string]int = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var part2GameToShape map[string]string = map[string]string{
	"A X": "Z", // Scissors
	"B X": "X", // Rock
	"C X": "Y", // Paper
	"A Y": "X", // Rock
	"B Y": "Y", // Paper
	"C Y": "Z", // Scissors
	"A Z": "Y", // Paper
	"B Z": "Z", // Scissors
	"C Z": "X", // Rock
}

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}
	inputSplit := strings.Split(string(bytes), "\n")

	var runningPoints int = 0

	// Part 1

	for _, game := range inputSplit {
		players := strings.Split(game, " ")
		runningPoints += shapePoints[players[1]]
		runningPoints += outcomePoints[game]
	}

	fmt.Println(runningPoints)

	// Part 2
	runningPoints = 0
	for _, game := range inputSplit {
		players := strings.Split(game, " ")
		runningPoints += shapePoints[part2GameToShape[game]]
		runningPoints += part2OutcomePoints[players[1]]
	}

	fmt.Println(runningPoints)
}
