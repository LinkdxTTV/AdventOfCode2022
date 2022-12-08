package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Part 1 is 4, Part 2 is 14
const lengthOfMarker = 14

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")
	letters := inputSplit[0]

	var counter int
	letterSet := map[string]int{}
	for i := 0; i < len(letters); i++ {
		if i >= lengthOfMarker {
			letterSet[string(letters[i-lengthOfMarker])]--
			if letterSet[string(letters[i-lengthOfMarker])] == 0 {
				delete(letterSet, string(letters[i-lengthOfMarker]))
			}
		}
		letterSet[string(letters[i])]++

		if len(letterSet) == lengthOfMarker {
			counter = i
			fmt.Println(letterSet)
			break
		}
	}
	// we off by 1 since we start at 0
	fmt.Println(counter + 1)
}
