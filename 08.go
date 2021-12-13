package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day08() {
	f, err := os.ReadFile("input/08.txt")
	ch(err)

	fmt.Println("Part 1:", CountDecode(string(f)))
}

func CountDecode(input string) int {
	count := 0
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)

	outputs := make([]string, 0)

	for s.Scan() {
		digits := strings.Split(s.Text(), " | ")[1]
		for _, v := range strings.Split(digits, " ") {
			outputs = append(outputs, v)
		}
	}

	for _, v := range outputs {
		switch len(v) {
		case 2, 4, 3, 7:
			count++
		}
	}

	return count
}
