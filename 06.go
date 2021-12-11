package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fishDays1 = 80
const fishDays2 = 256

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

	fish1 := NewLanternfish(string(b))

	for d := 1; d <= fishDays1; d++ {
		for _, f := range fish1 {
			maybeFish := f.LiveADay()
			if maybeFish != nil {
				fish1 = append(fish1, maybeFish)
			}
		}
	}

	fmt.Println("Part 1:", len(fish1))

	fish2 := NewLanternfish(string(b))

	for d := 1; d <= fishDays2; d++ {
		for _, f := range fish2 {
			maybeFish := f.LiveADay()
			if maybeFish != nil {
				fish2 = append(fish2, maybeFish)
			}
		}
	}

	fmt.Println("Part 2:", len(fish2))
}
