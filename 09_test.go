package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

const testinput09 = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestLowPoints(t *testing.T) {
	want := map[Point]int{
		{0, 1, 1}: 1,
		{0, 9, 0}: 0,
		{2, 2, 5}: 5,
		{4, 6, 5}: 5,
	}
	got := LowPoints(testinput09)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestBasins(t *testing.T) {
	want := 1134

	basins := Basins(testinput09)

	for _, b := range basins {
		fmt.Println(b)
	}

	sort.Slice(basins, func(i1, i2 int) bool {
		return len(basins[i1]) > len(basins[i2])
	})

	got := 1
	for _, b := range basins[:3] {
		// fmt.Println(len(b))
		got *= len(b)
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
