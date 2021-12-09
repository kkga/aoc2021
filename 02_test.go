package main

import "testing"

func TestApplyCommands(t *testing.T) {
	commands := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	want := Position{15, 60, 0}

	p := &Position{}
	p.ApplyCommands(commands, WithAim)

	if p.x != want.x || p.y != want.y {
		t.Errorf("want: %v, got: %v", want, p)
	}
}
