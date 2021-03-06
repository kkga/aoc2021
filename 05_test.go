package main

import (
	"reflect"
	"testing"
)

const testinput05 = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

// .......1..
// ..1....1..
// ..1....1..
// .......1..
// .112111211
// ..........
// ..........
// ..........
// ..........
// 222111....`

func TestNewDiagramWithoutDiagonals(t *testing.T) {
	want := Diagram{
		lines: []Line{
			{0, 9, 5, 9},
			{8, 0, 0, 8},
			{9, 4, 3, 4},
			{2, 2, 2, 1},
			{7, 0, 7, 4},
			{6, 4, 2, 0},
			{0, 9, 2, 9},
			{3, 4, 1, 4},
			{0, 0, 8, 8},
			{5, 5, 8, 2},
		},
		points: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
		},
	}
	got := NewDiagram(testinput05, WithoutDiagonals)

	// for _, p := range got.points {
	// 	fmt.Println(p)
	// }

	// log.Println(got.points)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNewDiagramWithDiagonals(t *testing.T) {
	want := Diagram{
		lines: []Line{
			{0, 9, 5, 9},
			{8, 0, 0, 8},
			{9, 4, 3, 4},
			{2, 2, 2, 1},
			{7, 0, 7, 4},
			{6, 4, 2, 0},
			{0, 9, 2, 9},
			{3, 4, 1, 4},
			{0, 0, 8, 8},
			{5, 5, 8, 2},
		},
		points: [][]int{
			{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
			{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
			{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
			{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
			{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
			{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
		},
	}
	got := NewDiagram(testinput05, WithDiagonals)

	// for _, p := range got.points {
	// 	fmt.Println(p)
	// }

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}
