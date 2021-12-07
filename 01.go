package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncreases(input []int) int {
	var increaseCount int

	for i, val := range input {
		if i == 0 {
			continue
		}
		if val > input[i-1] {
			increaseCount += 1
		}
	}
	return increaseCount
}

func Day01() {
	input, err := os.Open("input/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var measurements []int
	var numOfTimesIncreases int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, measurement)
	}

	numOfTimesIncreases = countIncreases(measurements)

	fmt.Println(numOfTimesIncreases)
}
