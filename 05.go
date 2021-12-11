package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Diagram struct {
	lines  []Line
	points [][]int
}

func NewDiagram(input string) (dia *Diagram) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)

	var lines []Line

	for s.Scan() {
		line := s.Text()
		points := strings.Split(line, " -> ")
		xy1 := strings.Split(points[0], ",")
		xy2 := strings.Split(points[1], ",")

		x1, _ := strconv.Atoi(xy1[0])
		x2, _ := strconv.Atoi(xy2[0])
		y1, _ := strconv.Atoi(xy1[1])
		y2, _ := strconv.Atoi(xy2[1])

		l := Line{x1, y1, x2, y2}

		lines = append(lines, l)
	}

	var d = &Diagram{lines, [][]int{}}

	for y := 0; y <= d.MaxY(); y++ {
		var xx = make([]int, d.MaxX()+1)
		for x := 0; x <= d.MaxX(); x++ {
			xx[x] = 0
		}
		d.points = append(d.points, xx)
	}

	for _, l := range d.lines {
		if l.x1 == l.x2 {
			if l.y1 < l.y2 {
				for y := l.y1; y <= l.y2; y++ {
					d.AddPoint(l.x1, y)
				}
			} else {
				for y := l.y1; y >= l.y2; y-- {
					d.AddPoint(l.x1, y)
				}
			}
		} else if l.y1 == l.y2 {
			if l.x1 < l.x2 {
				for x := l.x1; x <= l.x2; x++ {
					d.AddPoint(x, l.y1)
				}
			} else {
				for x := l.x1; x >= l.x2; x-- {
					d.AddPoint(x, l.y1)
				}
			}
		}
	}

	return d
}

func (d *Diagram) AddPoint(x, y int) {
	d.points[y][x] += 1
}

func (d Diagram) MaxX() int {
	var maxX int
	for _, l := range d.lines {
		maxX = maxInt([]int{maxX, l.x1, l.x2})
	}
	return maxX
}

func (d Diagram) MaxY() int {
	var maxY int
	for _, l := range d.lines {
		maxY = maxInt([]int{maxY, l.y1, l.y2})
	}
	return maxY
}

var maxInt = func(ii []int) int {
	var m int
	for _, i := range ii {
		if i > m {
			m = i
		}
	}
	return m
}

func Day05() {
	b, err := ioutil.ReadFile("input/05.txt")
	ch(err)

	d := NewDiagram(string(b))

	var overlappingPoints int

	for _, p := range d.points {
		for _, pp := range p {
			if pp > 1 {
				overlappingPoints++
			}
		}
	}

	fmt.Println("Part 1:", overlappingPoints)
}
