package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CalcGammaRate(input []string) (binRate string, err error) {
	var sb strings.Builder

	for i := 0; i < len(input[0]); i++ {
		b, err := MostCommonBit(input, i)
		if err != nil {
			return "", err
		}
		s := strconv.Itoa(b)
		sb.WriteString(s)
	}

	// log.Println("bin gam:", sb.String())

	binRate = sb.String()

	return
}

func GammaToEpsilon(gamma string) (binRate string, err error) {
	var sb strings.Builder

	for _, v := range gamma {
		switch string(v) {
		case "0":
			sb.WriteString("1")
		case "1":
			sb.WriteString("0")
		}
	}

	// log.Println("bin eps:", sb.String())

	binRate = sb.String()
	return
}

func MostCommonBit(input []string, bitIndex int) (bit int, err error) {
	var bits []int

	for _, val := range input {
		bit, err := strconv.Atoi(string(val[bitIndex]))
		if err != nil {
			return 0, err
		}
		bits = append(bits, bit)
	}

	zeros := 0
	ones := 0

	for _, bit := range bits {
		switch bit {
		case 0:
			zeros += 1
		case 1:
			ones += 1
		}
	}

	if zeros > ones {
		bit = 0
	} else {
		bit = 1
	}

	return
}

func Day03() {
	f, err := os.Open("input/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var input []string

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	gamma, err := CalcGammaRate(input)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := GammaToEpsilon(gamma)
	if err != nil {
		log.Fatal(err)
	}

	gammaDecimal, err := binToDec(gamma)
	if err != nil {
		log.Fatal(err)
	}

	epsilonDecimal, err := binToDec(epsilon)
	if err != nil {
		log.Fatal(err)
	}

	consumption := gammaDecimal * epsilonDecimal

	fmt.Println("Part 1:", consumption)
}

func binToDec(bin string) (int64, error) {
	decimal, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
