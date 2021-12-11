package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input/06.txt
var input string

const fishDays1 = 80
const fishDays2 = 256

func Day06() {
	fmt.Println("Part 1:", liveDays(input, 80))
	fmt.Println("Part 2:", liveDays(input, 256))
}

func liveDays(input string, days int) int {
	state := make([]int, 9)

	for _, num := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(num)
		daysLeft := n
		state[daysLeft]++
	}

	for d := 0; d < days; d++ {
		save := state[0]
		for i := 0; i < len(state)-1; i++ {
			state[i] = state[i+1]
		}
		state[8] = save
		state[6] += save
	}

	var sum int
	for _, n := range state {
		sum += n
	}

	return sum
}

// func (f fishlife) LiveDays(days int) int {
// 	for day := 1; day <= days; days++ {
// 		oldFish := f
// 		f := make(fishlife)
// 		for k, v := range oldFish {
// 			if k == 0 {
// 				f[6] += v
// 				f[8] += v
// 				continue
// 			}
// 			f[k-1] += v
// 		}
// 		// fmt.Println(f)
// 	}

// 	count := 0
// 	for _, v := range f {
// 		count += v
// 	}
// 	return count

// }

// type Lanternfish struct {
// 	timer int
// }

// func NewLanternfish(input string) (fish []*Lanternfish) {
// 	var timers []int

// 	nums := strings.Split(input, ",")
// 	for _, n := range nums {
// 		num, _ := strconv.Atoi(n)
// 		timers = append(timers, num)
// 	}

// 	for _, t := range timers {
// 		f := &Lanternfish{t}
// 		fish = append(fish, f)
// 	}
// 	return
// }

// func (f *Lanternfish) LiveADay() (newfish *Lanternfish) {
// 	if f.timer == 0 {
// 		newfish = &Lanternfish{8}
// 		f.timer = 6
// 	} else {
// 		f.timer--
// 	}

// 	return
// }

// func Day06() {
// 	b, err := os.ReadFile("input/06.txt")
// 	ch(err)

// 	fish1 := NewLanternfish(string(b))

// 	for d := 1; d <= fishDays1; d++ {
// 		for _, f := range fish1 {
// 			maybeFish := f.LiveADay()
// 			if maybeFish != nil {
// 				fish1 = append(fish1, maybeFish)
// 			}
// 		}
// 	}

// 	fmt.Println("Part 1:", len(fish1))

// 	fish2 := NewLanternfish(string(b))

// 	for d := 1; d <= fishDays2; d++ {
// 		log.Println(d, len(fish2))
// 		for _, f := range fish2 {
// 			maybeFish := f.LiveADay()
// 			if maybeFish != nil {
// 				fish2 = append(fish2, maybeFish)
// 			}
// 		}
// 	}

// 	fmt.Println("Part 2:", len(fish2))
// }
