package main

import (
	"reflect"
	"testing"
)

const testinput04 = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

func TestNewBingoBoard(t *testing.T) {
	input := `
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`
	want := BingoBoard{[]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}}
	got := NewBingoBoard(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestReadBingoNumOrder(t *testing.T) {
	want := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	got := ReadBingoNumOrder(testinput04)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestBingoMarkNumber(t *testing.T) {
	nums := []int{7, 10, 3, 3, 22, 9}
	b := &BingoBoard{[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}}
	want := &BingoBoard{[]int{-7, 4, -9, 5, 11, 17, 23, 2, 0, 14, 21, 24, -10, 16, 13, 6, 15, 25, 12, -22, 18, 20, 8, 19, -3, 26, 1}}

	for _, n := range nums {
		b.MarkNumber(n)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("want %v, got %v", want, b)
	}
}

func TestBingoWins(t *testing.T) {
	var tests = []struct {
		nums []int
		want bool
	}{
		{[]int{8, 2, 23, 24, 4}, true},
		{[]int{0, 24, 19, 7, 5}, true},
		{[]int{1, 1, 9, 1, 5}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			b := BingoBoard{[]int{
				22, 13, 17, 11, 0,
				8, 2, 23, 4, 24,
				21, 9, 14, 16, 7,
				6, 10, 3, 18, 5,
				1, 12, 20, 15, 19,
			}}
			for _, n := range tt.nums {
				b.MarkNumber(n)
			}
			got := b.Wins()
			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}

}

// func TestBingoScore(t *testing.T) {
// 	b := BingoBoard{[]int{
// 		22, 13, 17, 11, 0,
// 		8, 2, 23, 4, 24,
// 		21, 9, 14, 16, 7,
// 		6, 10, 3, 18, 5,
// 		1, 12, 20, 15, 19,
// 	}}
// 	nums := []int{22, 13, 17, 11, 0}
// }

func TestReadBingoInput(t *testing.T) {
	want := []BingoBoard{
		{[]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}},
		{[]int{3, 15, 0, 2, 22, 9, 18, 13, 17, 5, 19, 8, 7, 25, 23, 20, 11, 10, 24, 4, 14, 21, 16, 12, 6}},
		{[]int{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}},
	}

	t.Run("", func(t *testing.T) {
		got := ReadBingoInput(testinput04)
		for i, b := range got {
			if !reflect.DeepEqual(want[i], *b) {
				t.Errorf("want %v, got %v", want, got)
			}
		}
	})
}

func TestBingoRows(t *testing.T) {
	input := `
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
 `
	want := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	b := NewBingoBoard(input)

	t.Run("", func(t *testing.T) {
		got := b.Rows()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestBingoCols(t *testing.T) {
	input := `
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
 `
	want := [][]int{
		{22, 8, 21, 6, 1},
		{13, 2, 9, 10, 12},
		{17, 23, 14, 3, 20},
		{11, 4, 16, 18, 15},
		{0, 24, 7, 5, 19},
	}

	b := NewBingoBoard(input)

	t.Run("", func(t *testing.T) {
		got := b.Cols()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
