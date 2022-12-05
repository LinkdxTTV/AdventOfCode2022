package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	fullyContained := 0
	anyOverlap := 0

	for _, line := range inputSplit {
		elves := strings.Split(line, ",")
		leftElf := strings.Split(elves[0], "-")
		rightElf := strings.Split(elves[1], "-")
		leftElfLower, _ := strconv.Atoi(leftElf[0])
		leftElfHigher, _ := strconv.Atoi(leftElf[1])
		rightElfLower, _ := strconv.Atoi(rightElf[0])
		rightElfHigher, _ := strconv.Atoi(rightElf[1])

		if (leftElfLower <= rightElfLower && leftElfHigher >= rightElfHigher) || (rightElfLower <= leftElfLower && rightElfHigher >= leftElfHigher) {
			fullyContained++
		}

		if (leftElfLower >= rightElfLower && leftElfLower <= rightElfHigher) || (leftElfHigher >= rightElfLower && leftElfHigher <= rightElfHigher) || (rightElfLower >= leftElfLower && rightElfLower <= leftElfHigher) || (rightElfHigher >= leftElfLower && rightElfHigher <= leftElfHigher) {
			anyOverlap++
		}
	}
	// Hate overlap problems zz
	fmt.Println(fullyContained)
	fmt.Println(anyOverlap)
}
