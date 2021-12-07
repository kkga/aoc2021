package main

import "testing"

func TestCalculateIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7

	got := calculateIncreases(input)

	if got != want {
		t.Errorf("calculateIncreases: want %d, got %d", want, got)
	}

}
