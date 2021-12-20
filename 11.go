package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type octogrid [][]int

type position struct {
	row int
	col int
}

func Day11() {
	input, err := os.ReadFile("input/11.txt")
	ch(err)

	var flashes int
	o := newOctogrid(string(input))
	for i := 1; i <= 100; i++ {
		flashes += step(o)
	}

	fmt.Println("Part 1:", flashes)

	o = newOctogrid(string(input))
	flashsyncStep := 1
	for ; ; flashsyncStep++ {
		_ = step(o)
		if o.isFlashsync() {
			break
		}
	}
	fmt.Println("Part 2:", flashsyncStep)
}

func newOctogrid(input string) octogrid {
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Split(bufio.ScanLines)

	var octs [][]int

	for sc.Scan() {
		row := make([]int, len(sc.Text()))
		for i, r := range sc.Text() {
			oct, err := strconv.Atoi(string(r))
			ch(err)
			row[i] = oct
		}
		octs = append(octs, row)
	}

	return octs
}

func step(grid [][]int) (flashes int) {
	var toFlash []position

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 9 {
				toFlash = append(toFlash, position{row, col})
			}
			grid[row][col]++
		}
	}

	for ; len(toFlash) > 0; flashes++ {
		flashPos := toFlash[0]
		toFlash = toFlash[1:]
		ns := diagonalNeighbors(grid, flashPos)
		for _, n := range ns {
			if grid[n.row][n.col] == 9 {
				toFlash = append(toFlash, n)
			}
			grid[n.row][n.col]++
		}
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] > 9 {
				grid[row][col] = 0
			}
		}
	}

	return
}

func (o octogrid) String() string {
	var sb strings.Builder

	for row := range o {
		for col := range o[row] {
			sb.WriteString(fmt.Sprint(o[row][col]))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.String()
}

func (o octogrid) isFlashsync() bool {
	for row := range o {
		for col := range o[row] {
			if o[row][col] != 0 {
				return false
			}
		}
	}
	return true
}

func diagonalNeighbors(grid [][]int, p position) (neighbors []position) {
	maxCol := len(grid[0]) - 1
	maxRow := len(grid) - 1

	if p.row > 0 && p.col > 0 {
		neighbors = append(neighbors, position{p.row - 1, p.col - 1})
	}
	if p.row > 0 {
		neighbors = append(neighbors, position{p.row - 1, p.col})
	}
	if p.row > 0 && p.col < maxCol {
		neighbors = append(neighbors, position{p.row - 1, p.col + 1})
	}

	if p.col > 0 {
		neighbors = append(neighbors, position{p.row, p.col - 1})
	}
	if p.col < maxCol {
		neighbors = append(neighbors, position{p.row, p.col + 1})
	}

	if p.col > 0 && p.row < maxRow {
		neighbors = append(neighbors, position{p.row + 1, p.col - 1})
	}
	if p.row < maxRow {
		neighbors = append(neighbors, position{p.row + 1, p.col})
	}
	if p.row < maxRow && p.col < maxCol {
		neighbors = append(neighbors, position{p.row + 1, p.col + 1})
	}

	return
}
