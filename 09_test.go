package main

import (
	"reflect"
	"testing"
)

const testinput09 = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestLowPoints(t *testing.T) {
	want := []int{1, 0, 5, 5}
	got := LowPoints(testinput09)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
