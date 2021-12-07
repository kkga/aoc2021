package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncreases(input []int) (count int) {
	for i, val := range input {
		if i == 0 {
			continue
		}
		if val > input[i-1] {
			count += 1
		}
	}
	return
}

func intSliceToWindows(input []int, windowSize int) (windows [][]int) {
	for i := range input {
		end := i + windowSize

		if end > len(input) {
			break
		}

		windows = append(windows, input[i:end])
	}
	return
}

func countIncreasesInSumsOfWindows(input []int, windowSize int) (count int) {
	var sums []int

	windows := intSliceToWindows(input, windowSize)

	for _, window := range windows {
		var sum int
		for _, num := range window {
			sum += num
		}
		sums = append(sums, sum)
	}

	count = countIncreases(sums)
	return
}

func Day01() {
	input, err := os.Open("input/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var measurements []int
	var numOfTimesIncreases int
	var numOfTimesIncreasesInWindowsOf3 int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, measurement)
	}

	numOfTimesIncreases = countIncreases(measurements)
	numOfTimesIncreasesInWindowsOf3 = countIncreasesInSumsOfWindows(measurements, 3)

	fmt.Println("part 1:", numOfTimesIncreases)
	fmt.Println("part 2:", numOfTimesIncreasesInWindowsOf3)
}
