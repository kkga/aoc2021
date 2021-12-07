package main

import "testing"

func TestCountIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7

	got := countIncreases(input)

	if got != want {
		t.Errorf("calculateIncreases: want %d, got %d", want, got)
	}

}
