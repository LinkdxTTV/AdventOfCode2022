package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	duplicateLetterList := []string{}
	priorityRunningSum := 0

	for _, rucksack := range inputSplit {
		left, right := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		var duplicateLetter string
		leftLetters := map[string]bool{}
		for _, char := range left {
			leftLetters[string(char)] = true
		}

		for _, char := range right {
			_, ok := leftLetters[string(char)]
			if ok {
				duplicateLetter = string(char)
				break
			}
		}
		if duplicateLetter == "" {
			log.Fatal("No duplicate letter?", rucksack, left, right)
		}

		duplicateLetterList = append(duplicateLetterList, duplicateLetter)
		priorityRunningSum += turnASCIItoDEC(duplicateLetter)

	}

	fmt.Println(priorityRunningSum)

	// Part 2

	group := []string{}
	groupLetters := []string{}
	groupSum := 0

	for i := 0; i < len(inputSplit); i++ {
		group = append(group, inputSplit[i])
		if len(group) == 3 {
			// Do stuff

			first := map[string]bool{}
			for _, char := range group[0] {
				first[string(char)] = true
			}
			second := map[string]bool{}
			for _, char := range group[1] {
				if first[string(char)] {
					second[string(char)] = true
				}
			}
			for _, char := range group[2] {
				if second[string(char)] {
					groupLetters = append(groupLetters, string(char))
					groupSum += turnASCIItoDEC(string(char))

					break
				}
			}

			// Reset
			group = []string{}
		}
	}
	fmt.Println(groupSum)
}

func turnASCIItoDEC(input string) int {
	// Unlucky
	uppercase := strings.ToUpper(input)
	if input == uppercase {
		return int(input[0]) - 64 + 26
	}

	return int(input[0]) - 96
}
