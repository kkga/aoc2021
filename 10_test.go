package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

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

func TestIncompleteChars(t *testing.T) {
	want1 := 26397
	want2 := 288957

	sc := bufio.NewScanner(strings.NewReader(testinput10))
	sc.Split(bufio.ScanLines)

	lines := make(map[string][]rune)
	corrupted := make(map[string]rune)

	for sc.Scan() {
		stack, err := IncompleteChars(sc.Text())
		if err != nil {
			corrupted[sc.Text()] = stack[0]
		} else {
			lines[sc.Text()] = stack
		}
	}

	p1score := 0
	for _, v := range corrupted {
		p1score += scoreMap[v]
	}

	if p1score != want1 {
		t.Fatalf("want %d, got %d", want1, p1score)
	}

	p2scores := make([]int, 0)
	for _, v := range lines {
		score := 0
		rev := reverse(v)
		for _, ch := range rev {
			score *= 5
			score += completionScoreMap[charMap[ch]]
		}
		p2scores = append(p2scores, score)
	}
	sort.Ints(p2scores)

	p2score := p2scores[len(p2scores)/2]

	if p2score != want2 {
		t.Fatalf("want %d, got %d", want2, p2score)
	}
}
