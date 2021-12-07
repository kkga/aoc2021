package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x   int
	y   int
	aim int
}

func (p *Position) ApplyCommands(commands []string) error {
	for _, command := range commands {
		dir := strings.Split(command, " ")[0]
		length, err := strconv.Atoi(strings.Split(command, " ")[1])
		if err != nil {
			return err
		}

		switch dir {
		case "down":
			p.aim += length
		case "up":
			p.aim -= length
		case "forward":
			p.x += length
			p.y += p.aim * length
		}

	}
	return nil
}

func (p *Position) Multiplied() int {
	return p.x * p.y
}

func (p *Position) String() string {
	return fmt.Sprintf("horizontal: %d, depth: %d", p.x, p.y)
}

func Day02() {
	input, err := os.Open("input/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var commands []string

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	p := &Position{}
	p.ApplyCommands(commands)

	multiplied := p.Multiplied()

	fmt.Println("Part 1:", multiplied)

}
