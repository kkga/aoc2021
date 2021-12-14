package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a day number in XX format")
		os.Exit(1)
	}

	day := os.Args[1]

	switch day {
	case "01":
		Day01()
	case "02":
		Day02()
	case "03":
		Day03()
	case "04":
		Day04()
	case "05":
		Day05()
	case "06":
		Day06()
	case "07":
		Day07()
	case "08":
		Day08()
	case "09":
		Day09()
	case "10":
		Day10()
	default:
		log.Fatalf("There's no solution for day %s", day)
	}
}

func ch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
