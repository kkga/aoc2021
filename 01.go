package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	for i, val := range measurements {
		if i == 0 {
			continue
		}
		if val > measurements[i-1] {
			numOfTimesIncreases += 1
		}
	}

	fmt.Println(numOfTimesIncreases)
}
