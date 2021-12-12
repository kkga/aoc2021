package main

import "testing"

const testinput07 = "16,1,2,0,4,2,7,1,2,14"

func TestCheapestXMove(t *testing.T) {
	want := 37
	got := CheapestXMove(testinput07, WithoutIncrease)

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCheapestXMoveWithIncrease(t *testing.T) {
	want := 168
	got := CheapestXMove(testinput07, WithIncrease)

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
