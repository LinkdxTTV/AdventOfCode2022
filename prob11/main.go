package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	inspects      int
	worries       []int
	worryModifier func([]*Monkey, int)
}

func main() {
	monkey0 := &Monkey{
		inspects: 0,
		worries:  []int{63, 57},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry * 11
			worry = relief(worry)
			if worry%7 == 0 {
				monkeys[6].appendToWorries(worry)
			} else {
				monkeys[2].appendToWorries(worry)
			}
		},
	}

	monkey1 := &Monkey{
		inspects: 0,
		worries:  []int{82, 66, 87, 78, 77, 92, 83},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry + 1
			worry = relief(worry)
			if worry%11 == 0 {
				monkeys[5].appendToWorries(worry)
			} else {
				monkeys[0].appendToWorries(worry)
			}
		},
	}

	monkey2 := &Monkey{
		inspects: 0,
		worries:  []int{97, 53, 53, 85, 58, 54},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry * 7
			worry = relief(worry)
			if worry%13 == 0 {
				monkeys[4].appendToWorries(worry)
			} else {
				monkeys[3].appendToWorries(worry)
			}
		},
	}

	monkey3 := &Monkey{
		inspects: 0,
		worries:  []int{50},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry + 3
			worry = relief(worry)
			if worry%3 == 0 {
				monkeys[1].appendToWorries(worry)
			} else {
				monkeys[7].appendToWorries(worry)
			}
		},
	}

	monkey4 := &Monkey{
		inspects: 0,
		worries:  []int{64, 69, 52, 65, 73},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry + 6
			worry = relief(worry)
			if worry%17 == 0 {
				monkeys[3].appendToWorries(worry)
			} else {
				monkeys[7].appendToWorries(worry)
			}
		},
	}

	monkey5 := &Monkey{
		inspects: 0,
		worries:  []int{57, 91, 65},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry + 5
			worry = relief(worry)
			if worry%2 == 0 {
				monkeys[0].appendToWorries(worry)
			} else {
				monkeys[6].appendToWorries(worry)
			}
		},
	}

	monkey6 := &Monkey{
		inspects: 0,
		worries:  []int{67, 91, 84, 78, 60, 69, 99, 83},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry * worry
			worry = relief(worry)
			if worry%5 == 0 {
				monkeys[2].appendToWorries(worry)
			} else {
				monkeys[4].appendToWorries(worry)
			}
		},
	}

	monkey7 := &Monkey{
		inspects: 0,
		worries:  []int{58, 78, 69, 65},
		worryModifier: func(monkeys []*Monkey, worry int) {
			worry = worry + 7
			worry = relief(worry)
			if worry%19 == 0 {
				monkeys[5].appendToWorries(worry)
			} else {
				monkeys[1].appendToWorries(worry)
			}
		},
	}

	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			monkey.inspect(monkeys)
		}
	}

	inspections := []int{}
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspects)
	}

	sort.Ints(inspections)
	fmt.Println(inspections)
	lenghty := len(inspections)
	fmt.Println(inspections[lenghty-1] * inspections[lenghty-2])
}

func (m *Monkey) appendToWorries(worry int) {
	m.worries = append(m.worries, worry)
}

func (m *Monkey) inspect(monkeys []*Monkey) {
	for _, worry := range m.worries {
		m.inspects++
		m.worryModifier(monkeys, worry)
	}
	m.worries = []int{}
}

// MATH POG?
// This 9699690 is the product of all the prime numbers above that the monkeys test for

// Part 1 relief uses worry / 3
func relief(worry int) int {
	for worry > 9699690 {
		num := worry / 9699690
		worry += -9699690 * num
	}
	return worry
}
