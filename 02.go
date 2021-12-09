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

type CommandOption int

const (
	WithAim CommandOption = iota
	WithoutAim
)

func (p *Position) ApplyCommands(commands []string, opt CommandOption) error {
	for _, command := range commands {
		dir := strings.Split(command, " ")[0]
		length, err := strconv.Atoi(strings.Split(command, " ")[1])
		if err != nil {
			return err
		}

		if opt == WithAim {
			switch dir {
			case "down":
				p.aim += length
			case "up":
				p.aim -= length
			case "forward":
				p.x += length
				p.y += p.aim * length
			}
		} else {
			switch dir {
			case "down":
				p.y += length
			case "up":
				p.y -= length
			case "forward":
				p.x += length
			}
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

	p1 := &Position{}
	p1.ApplyCommands(commands, WithoutAim)
	p1Multiplied := p1.Multiplied()

	p2 := &Position{}
	p2.ApplyCommands(commands, WithAim)
	p2Multiplied := p2.Multiplied()

	fmt.Println("Part 1:", p1Multiplied)
	fmt.Println("Part 2:", p2Multiplied)

}
