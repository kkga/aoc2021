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

	fuel1 := CheapestXMove(string(b), WithoutIncrease)
	fmt.Println("Part 1:", fuel1)

	fuel2 := CheapestXMove(string(b), WithIncrease)
	fmt.Println("Part 2:", fuel2)
}

type XMoveOption int

const (
	WithIncrease XMoveOption = iota
	WithoutIncrease
)

func CheapestXMove(input string, opt XMoveOption) int {
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
			diff := int(math.Abs(float64(crab - x)))
			cost += calcFuel(diff, opt)
		}
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func calcFuel(diff int, opt XMoveOption) int {
	if opt == WithoutIncrease {
		return diff
	} else {
		total := 0
		for i := 1; i <= diff; i++ {
			total += i
		}
		return total
	}
}
