package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCommonBit(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	var tests = []struct {
		criteria BitCriteria
		bitIdx   int
		want     int
	}{
		{MostCommon, 0, 1},
		{MostCommon, 1, 0},
		{MostCommon, 2, 1},
		{MostCommon, 3, 1},
		{MostCommon, 4, 0},
		{LeastCommon, 0, 0},
		{LeastCommon, 1, 1},
		{LeastCommon, 2, 0},
		{LeastCommon, 3, 0},
		{LeastCommon, 4, 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("bit: %d", tt.bitIdx), func(t *testing.T) {
			got, _ := CommonBit(input, tt.bitIdx, tt.criteria)
			if got != tt.want {
				t.Errorf("bit %d failed: want: %d, got: %d", tt.bitIdx, tt.want, got)
			}
		})
	}
}

func TestFilterByCriteria(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	var tests = []struct {
		criteria BitCriteria
		want     []string
	}{
		{LeastCommon, []string{"01010"}},
		{MostCommon, []string{"10111"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, _ := FilterByCriteria(input, tt.criteria)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestGammaRate(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	var want int64 = 22

	gamma, err := GammaRate(input)
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

	gamma, err := GammaRate(input)
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
