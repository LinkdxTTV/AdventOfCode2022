package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	point     point
	visited   map[point]bool
	following *Segment
}

type point struct {
	x int
	y int
}

type Turn struct {
	direction string
	value     int
}

func main() {
	bytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal("Could not read input")
	}

	inputSplit := strings.Split(string(bytes), "\n")

	turns := []Turn{}

	for _, line := range inputSplit {
		if line == "" {
			continue
		}
		strSplit := strings.Split(line, " ")
		direction := strSplit[0]
		amount, err := strconv.Atoi(strSplit[1])
		if err != nil {
			log.Fatal("bad int parse", strSplit)
		}
		turns = append(turns, Turn{direction, amount})
	}

	head := Segment{point{0, 0}, map[point]bool{{0, 0}: true}, nil}
	tail := Segment{point{0, 0}, map[point]bool{{0, 0}: true}, &head}

	for _, turn := range turns {
		for i := 0; i < turn.value; i++ {
			xDiff := 0
			yDiff := 0
			switch turn.direction {
			case "U":
				yDiff = 1
			case "D":
				yDiff = -1
			case "L":
				xDiff = -1
			case "R":
				xDiff = 1
			}
			head.point.move(xDiff, yDiff)
			tail.catchUpSegmentIfNeeded()
		}
	}

	fmt.Println("Part 1: ", len(tail.visited))

	// Part 2

	helpfulSegmentMap := map[int]*Segment{}

	head = Segment{point{0, 0}, map[point]bool{{0, 0}: true}, nil}
	helpfulSegmentMap[0] = &head
	for i := 1; i <= 9; i++ {
		newSegment := Segment{point{0, 0}, map[point]bool{{0, 0}: true}, helpfulSegmentMap[i-1]}
		helpfulSegmentMap[i] = &newSegment
	}

	head = *helpfulSegmentMap[0]

	for _, turn := range turns {
		for i := 0; i < turn.value; i++ {
			xDiff := 0
			yDiff := 0
			switch turn.direction {
			case "U":
				yDiff = 1
			case "D":
				yDiff = -1
			case "L":
				xDiff = -1
			case "R":
				xDiff = 1
			}
			head.point.move(xDiff, yDiff)
			for i := 1; i <= 9; i++ {
				helpfulSegmentMap[i].catchUpSegmentIfNeeded()
			}
		}
	}

	fmt.Println("Part 2: ", len(helpfulSegmentMap[9].visited))
}

func (t *Segment) catchUpSegmentIfNeeded() {
	if t.isTouchingHead() {
		return
	}
	xDiff := t.following.point.x - t.point.x
	yDiff := t.following.point.y - t.point.y
	if xDiff >= 2 {
		xDiff = 1
	}
	if xDiff <= -2 {
		xDiff = -1
	}
	if yDiff >= 2 {
		yDiff = 1
	}
	if yDiff <= -2 {
		yDiff = -1
	}

	t.point.move(xDiff, yDiff)
	t.visited[t.point] = true
}

func (t *Segment) isTouchingHead() bool {
	if abs(t.following.point.x-t.point.x) > 1 || abs(t.following.point.y-t.point.y) > 1 {
		return false
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p *point) move(xDiff, yDiff int) {
	p.x += xDiff
	p.y += yDiff
}
