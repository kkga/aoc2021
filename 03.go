package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BitCriteria int

const (
	LeastCommon BitCriteria = iota
	MostCommon
)

func FilterByCriteria(input []string, criteria BitCriteria) (string, error) {
	filtered := input

	for i := 0; i < len(input[0]); i++ {
		if len(filtered) == 1 {
			break
		}

		cb, err := CommonBit(filtered, i, criteria)
		if err != nil {
			return "", err
		}

		var n = make([]string, 0)

		for _, val := range filtered {
			b, err := strconv.Atoi(string(val[i]))
			if err != nil {
				return "", err
			}
			if b == cb {
				n = append(n, val)
			}
		}

		filtered = n
	}

	return filtered[0], nil
}

func CommonBit(input []string, bitIndex int, criteria BitCriteria) (bit int, err error) {
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

	switch criteria {
	case MostCommon:
		if zeros > ones {
			bit = 0
		} else {
			bit = 1
		}
	case LeastCommon:
		if zeros > ones {
			bit = 1
		} else {
			bit = 0
		}
	}
	return
}

func GammaRate(input []string) (binRate string, err error) {
	var sb strings.Builder

	for i := 0; i < len(input[0]); i++ {
		b, err := CommonBit(input, i, MostCommon)
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

func EpsilonRate(input []string) (binRate string, err error) {
	var sb strings.Builder

	for i := 0; i < len(input[0]); i++ {
		b, err := CommonBit(input, i, LeastCommon)
		if err != nil {
			return "", err
		}
		s := strconv.Itoa(b)
		sb.WriteString(s)
	}

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

	gamma, err := GammaRate(input)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := EpsilonRate(input)
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

	o2, err := FilterByCriteria(input, MostCommon)
	if err != nil {
		log.Fatal(err)
	}
	o2decimal, err := binToDec(o2)
	if err != nil {
		log.Fatal(err)
	}

	co2, err := FilterByCriteria(input, LeastCommon)
	if err != nil {
		log.Fatal(err)
	}
	co2decimal, err := binToDec(co2)
	if err != nil {
		log.Fatal(err)
	}

	lifeSupport := o2decimal * co2decimal

	fmt.Println("Part 1:", consumption)
	fmt.Println("Part 2:", lifeSupport)
}

func binToDec(bin string) (int64, error) {
	decimal, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
