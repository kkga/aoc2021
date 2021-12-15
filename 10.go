package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScoreMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var charMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Day10() {
	input, err := os.ReadFile("input/10.txt")
	ch(err)

	lines := make(map[string][]rune)
	corrupted := make(map[string]rune)

	sc := bufio.NewScanner(strings.NewReader(string(input)))
	sc.Split(bufio.ScanLines)
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
	fmt.Println("Part 1:", p1score)

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
	fmt.Println("Part 2:", p2score)
}

func IncompleteChars(line string) (stack []rune, err error) {
	stack = make([]rune, 0)

Out:
	for _, ch := range line {
		switch ch {
		case '(', '[', '{', '<':
			stack = append(stack, ch)
		case ')', ']', '}', '>':
			expected := charMap[stack[len(stack)-1]]
			if ch != expected {
				// fmt.Println(line)
				// fmt.Printf("expected: %s, found: %s\n", string(expected), string(ch))
				stack = []rune{ch}
				err = errors.New("corrupted line")
				break Out
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return
}

func reverse(s []rune) []rune {
	a := make([]rune, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
