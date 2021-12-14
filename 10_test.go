package main

import "testing"

const testinput10 = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestCheckMatchingChars(t *testing.T) {
	want := 26397
	got := CheckMatchingChars(testinput10)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
