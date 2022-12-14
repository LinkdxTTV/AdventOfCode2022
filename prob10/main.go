package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type register struct {
	value int
}

type command struct {
	add int
}

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	commands := []command{}

	for _, text := range inputSplit {
		if text == "noop" {
			commands = append(commands, command{0})
			continue
		}
		val := strings.Split(text, " ")
		addedValue, err := strconv.Atoi(val[1])
		if err != nil {
			log.Fatal("parse error")
		}
		commands = append(commands, []command{{0}, {addedValue}}...)
	}

	scoreMap := map[int]int{}
	registerMap := map[int]int{}

	xRegister := register{1}
	for i, nextCommand := range commands {
		cycle := i + 1
		scoreMap[cycle] = cycle * xRegister.value
		registerMap[cycle] = xRegister.value

		xRegister.applyCommand(nextCommand)

	}

	sum := scoreMap[20] + scoreMap[60] + scoreMap[100] + scoreMap[140] + scoreMap[180] + scoreMap[220]
	fmt.Println(sum)

	for i := 1; i <= 240; i++ {
		if (i-1)%40 == 0 {
			fmt.Println()
		}
		if compareSpriteToCycle((i-1)%40, registerMap[i]) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (r *register) applyCommand(com command) {
	r.value += com.add
}

func compareSpriteToCycle(i, value int) bool {
	return (abs(value-i) <= 1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
