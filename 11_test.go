package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testinput11 = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestStep(t *testing.T) {
	o := newOctogrid(testinput11)

	want10step := 204
	var got10step int
	want100step := 1656
	var got100step int

	for i := 1; i <= 10; i++ {
		got10step += step(o)
		// if i%10 == 0 {
		// 	fmt.Println(i)
		// 	fmt.Println(o)
		// }
	}

	o = newOctogrid(testinput11)

	for i := 1; i <= 100; i++ {
		got100step += step(o)
		// if i%10 == 0 {
		// 	fmt.Println(i)
		// 	fmt.Println(o)
		// }
	}

	if !cmp.Equal(want10step, got10step) {
		t.Error(cmp.Diff(want10step, got10step))
	}

	if !cmp.Equal(want100step, got100step) {
		t.Error(cmp.Diff(want100step, got100step))
	}

}
