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
	default:
		log.Fatalf("There's no solution for day %s", day)
	}
}
