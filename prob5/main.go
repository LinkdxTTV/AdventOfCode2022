package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Move struct {
	amount int
	from   int
	to     int
}

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	// Initialize slices

	stacks := map[int][]string{}

	// Just hardcode it
	stacks[1] = []string{"F", "T", "C", "L", "R", "P", "G", "Q"}
	stacks[2] = []string{"N", "Q", "H", "W", "R", "F", "S", "J"}
	stacks[3] = []string{"F", "B", "H", "W", "P", "M", "Q"}
	stacks[4] = []string{"V", "S", "T", "D", "F"}
	stacks[5] = []string{"Q", "L", "D", "W", "V", "F", "Z"}
	stacks[6] = []string{"Z", "C", "L", "S"}
	stacks[7] = []string{"Z", "B", "M", "V", "D", "F"}
	stacks[8] = []string{"T", "J", "B"}
	stacks[9] = []string{"Q", "N", "B", "G", "L", "S", "P", "H"}

	regex := regexp.MustCompile("[0-9]+")
	moves := []Move{}

	for _, line := range inputSplit {
		if !strings.HasPrefix(line, "move") {
			continue
		}

		nums := regex.FindAllString(line, -1)
		moves = append(moves, turnRegexIntoMove(nums))
	}

	for _, move := range moves {
		fromStack := stacks[move.from]
		toStack := stacks[move.to]

		var crane []string

		crane, fromStack = fromStack[(len(fromStack)-move.amount):], fromStack[:(len(fromStack)-move.amount)]

		/* Part 1
		toStack = append(toStack, reverseSlice(crane)...)
		*/

		// Part 2
		toStack = append(toStack, (crane)...)
		//
		stacks[move.from], stacks[move.to] = fromStack, toStack
	}

	fmt.Println(printTopOfStack(stacks))
}

func turnRegexIntoMove(input []string) Move {
	move := Move{}
	move.amount, _ = strconv.Atoi(input[0])
	move.from, _ = strconv.Atoi(input[1])
	move.to, _ = strconv.Atoi(input[2])

	return move
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func printTopOfStack(stacks map[int][]string) string {
	output := ""
	for i := 1; i <= 9; i++ {
		output += (stacks[i][len(stacks[i])-1])
	}
	return output
}
