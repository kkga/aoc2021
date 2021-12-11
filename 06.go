package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fishDays = 80

type Lanternfish struct {
	timer int
}

func NewLanternfish(input string) (fish []*Lanternfish) {
	var timers []int

	nums := strings.Split(input, ",")
	for _, n := range nums {
		num, _ := strconv.Atoi(n)
		timers = append(timers, num)
	}

	for _, t := range timers {
		f := &Lanternfish{t}
		fish = append(fish, f)
	}
	return
}

func (f *Lanternfish) LiveADay() (newfish *Lanternfish) {
	if f.timer == 0 {
		newfish = &Lanternfish{8}
		f.timer = 6
	} else {
		f.timer--
	}

	return
}

func Day06() {
	b, err := os.ReadFile("input/06.txt")
	ch(err)

	fish := NewLanternfish(string(b))

	for d := 1; d <= fishDays; d++ {
		for _, f := range fish {
			maybeFish := f.LiveADay()
			if maybeFish != nil {
				fish = append(fish, maybeFish)
			}
		}
	}

	fmt.Println("Part 1:", len(fish))
}
