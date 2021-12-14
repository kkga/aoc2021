package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	row int
	col int
	val int
}

func Day09() {
	input, err := os.ReadFile("input/09.txt")
	ch(err)

	lp := LowPoints(string(input))
	var risk int

	for _, p := range lp {
		risk += 1 + p.val
	}

	fmt.Println("Part 1:", risk)

	basins := Basins(string(input))

	sort.Slice(basins, func(i1, i2 int) bool {
		return len(basins[i1]) > len(basins[i2])
	})

	mlt := 1
	for _, b := range basins[:3] {
		mlt *= len(b)
	}

	fmt.Println("Part 2:", mlt)
}

func LowPoints(input string) (points []Point) {
	hmap := heightmap(input)

	for i, r := range hmap {
		for j, n := range r {
			p := Point{i, j, n}
			if isLowPoint(hmap, p) {
				points = append(points, p)
			}
		}
	}

	return
}

func Basins(input string) (basins [][]Point) {
	lp := LowPoints(input)
	hmap := heightmap(input)

	// Tried a recursive approach first, but something's wrong
	//
	// var basin func(b []Point, step int, ns []Point) []Point
	// basin = func(b []Point, step int, ns []Point) []Point {
	// 	if step == 9 {
	// 		return b
	// 	}
	// 	var nns []Point
	// 	for _, n := range ns {
	// 		if n.val == step {
	// 			// fmt.Println(step)
	// 			b = append(b, n)
	// 			for _, nn := range neighbors(hmap, n) {
	// 				if !hasPoint(nns, nn) {
	// 					nns = append(nns, nn)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return basin(b, step+1, nns)
	// }

	for _, p := range lp {
		basin := []Point{p}
		ns := neighbors(hmap, p)
		step := p.val + 1

		for len(ns) > 0 {
			if step == 9 {
				break
			}
			for _, n := range ns {
				if n.val == step {
					if !hasPoint(basin, n) {
						basin = append(basin, n)
					}
					ns = append(ns, neighbors(hmap, n)...)
				}
			}
			step++
		}

		basins = append(basins, basin)
	}

	return
}

func remove(slice []Point, s int) []Point {
	return append(slice[:s], slice[s+1:]...)
}

func hasPoint(points []Point, p Point) bool {
	for _, pp := range points {
		if pp.col == p.col && pp.row == p.row {
			return true
		}
	}
	return false
}

func heightmap(input string) (hmap [][]int) {
	rows := strings.Split(input, "\n")
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
	return
}

func neighbors(hmap [][]int, p Point) (neighbors []Point) {
	if p.row > 0 {
		neighbors = append(neighbors, Point{p.row - 1, p.col, hmap[p.row-1][p.col]})
	}
	if p.col > 0 {
		neighbors = append(neighbors, Point{p.row, p.col - 1, hmap[p.row][p.col-1]})
	}
	if p.row < len(hmap)-1 {
		neighbors = append(neighbors, Point{p.row + 1, p.col, hmap[p.row+1][p.col]})
	}
	if p.col < len(hmap[0])-1 {
		neighbors = append(neighbors, Point{p.row, p.col + 1, hmap[p.row][p.col+1]})
	}
	return
}

func isLowPoint(hmap [][]int, p Point) bool {
	for _, point := range neighbors(hmap, p) {
		if point.val <= hmap[p.row][p.col] {
			return false
		}
	}

	return true
}
