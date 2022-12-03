package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")
	fmt.Println(inputSplit)

	sums := []int{}
	var runningSum int

	for _, line := range inputSplit {
		if line == "" {
			sums = append(sums, runningSum)
			runningSum = 0
			continue
		}
		num, err := strconv.Atoi(line)
		fmt.Println(num)
		if err != nil {
			log.Fatal(err, "Wtf is this?", line)
		}
		runningSum += num
	}

	sort.Ints(sums)
	fmt.Println(sums)

	sumsLen := len(sums)

	fmt.Println("Fattest Elf: ", sums[sumsLen-1])
	fmt.Println("Fattest 3 Elves: ", sums[sumsLen-1]+sums[sumsLen-2]+sums[sumsLen-3])
}
