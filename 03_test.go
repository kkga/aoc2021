package main

import (
	"fmt"
	"testing"
)

func TestMostCommonBit(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	var tests = []struct {
		bitIdx int
		want   int
	}{
		{0, 1},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("bit: %d", tt.bitIdx), func(t *testing.T) {
			got, _ := MostCommonBit(input, tt.bitIdx)
			if got != tt.want {
				t.Errorf("bit %d failed: want: %d, got: %d", tt.bitIdx, tt.want, got)
			}
		})
	}

}

func TestCalcGammaRate(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	var want int64 = 22

	gamma, err := CalcGammaRate(input)
	if err != nil {
		t.Error(err)
	}
	got, err := binToDec(gamma)
	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGammaToEpsilon(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	var want int64 = 9

	gamma, err := CalcGammaRate(input)
	if err != nil {
		t.Error(err)
	}
	epsilon, err := GammaToEpsilon(gamma)
	if err != nil {
		t.Error(err)
	}
	got, err := binToDec(epsilon)
	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
