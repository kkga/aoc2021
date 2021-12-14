package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day09() {
	input, err := os.ReadFile("input/09.txt")
	ch(err)

	lp := LowPoints(string(input))
	var risk int

	for _, p := range lp {
		risk += 1 + p
	}

	fmt.Println("Part 1:", risk)
}

func LowPoints(input string) (points []int) {
	rows := strings.Split(input, "\n")

	var hmap [][]int

	for _, r := range rows {
		if r == "" {
			continue
		}
		var row []int
		for _, n := range r {
			num, err := strconv.Atoi(string(n))
			ch(err)

			row = append(row, num)
		}
		hmap = append(hmap, row)
	}

	for i, r := range hmap {
		for j, n := range r {
			if isLowPoint(hmap, i, j) {
				points = append(points, n)
			}
		}
	}

	return
}

func isLowPoint(hmap [][]int, row int, col int) bool {
	var neighbors []int

	if row > 0 {
		neighbors = append(neighbors, hmap[row-1][col])
	}
	if col > 0 {
		neighbors = append(neighbors, hmap[row][col-1])
	}
	if row < len(hmap)-1 {
		neighbors = append(neighbors, hmap[row+1][col])
	}
	if col < len(hmap[0])-1 {
		neighbors = append(neighbors, hmap[row][col+1])
	}

	for _, num := range neighbors {
		if num <= hmap[row][col] {
			return false
		}
	}

	return true
}
