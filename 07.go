package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day07() {
	b, err := os.ReadFile("input/07.txt")
	ch(err)

	fuel := CheapestXMove(string(b))

	fmt.Println("Part 1:", fuel)
}

func CheapestXMove(input string) int {
	var crabs []int

	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		crab, err := strconv.Atoi(s)
		ch(err)
		crabs = append(crabs, crab)
	}

	maxX := crabs[0]
	minX := crabs[0]

	for _, c := range crabs {
		if c > maxX {
			maxX = c
		}
		if c < minX {
			minX = c
		}
	}

	minCost := math.MaxInt

	for x := minX; x <= maxX; x++ {
		var cost int
		for _, crab := range crabs {
			fuel := int(math.Abs(float64(crab - x)))
			cost += fuel
		}
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
