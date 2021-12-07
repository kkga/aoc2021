package main

import (
	"reflect"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7

	got := countIncreases(input)

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestIntSliceToWindows(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	windowSize := 3
	want := [][]int{
		{199, 200, 208},
		{200, 208, 210},
		{208, 210, 200},
		{210, 200, 207},
		{200, 207, 240},
		{207, 240, 269},
		{240, 269, 260},
		{269, 260, 263},
	}

	got := intSliceToWindows(input, windowSize)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCountIncreasesInSumsOfWindows(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5

	got := countIncreasesInSumsOfWindows(input, 3)

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
