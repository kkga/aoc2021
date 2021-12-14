package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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

	score := CheckMatchingChars(string(input))
	fmt.Println("Part 1:", score)
}

func CheckMatchingChars(input string) int {
	var score int

	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Split(bufio.ScanLines)

	charStack := make([]rune, 0)
	corrupted := make([]string, 0)
	illegalChars := make([]rune, 0)

	for sc.Scan() {
	Out:
		for _, ch := range sc.Text() {
			switch ch {
			case '(', '[', '{', '<':
				charStack = append(charStack, ch)
			default:
				if ch != charMap[charStack[len(charStack)-1]] {
					corrupted = append(corrupted, sc.Text())
					illegalChars = append(illegalChars, ch)
					break Out
				} else {
					charStack = charStack[:len(charStack)-1]
				}
			}
		}
	}

	// for _, l := range corrupted {
	// 	fmt.Println(l)
	// }

	for _, ch := range illegalChars {
		score += scoreMap[ch]
	}

	return score
}
