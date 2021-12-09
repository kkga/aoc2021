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

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func FilterByCriteria(input []string, criteria BitCriteria) ([]string, error) {
	filtered := input

	for i := 0; i < len(input[0]); i++ {
		cb, err := CommonBit(input, i, criteria)
		if err != nil {
			return nil, err
		}

		for j := 0; j < len(filtered); j++ {
			b, err := strconv.Atoi(string(filtered[j][i]))
			if err != nil {
				return nil, err
			}
			log.Println(filtered)

			if b != cb {
				filtered = remove(filtered, j)
			}
		}
	}

	return filtered, nil
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
	case LeastCommon:
		if zeros > ones {
			bit = 1
		} else {
			bit = 0
		}
	case MostCommon:
		if zeros > ones {
			bit = 0
		} else {
			bit = 1
		}
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

	fmt.Println("Part 1:", consumption)
}

func binToDec(bin string) (int64, error) {
	decimal, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}
