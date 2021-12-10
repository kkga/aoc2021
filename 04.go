package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const BingoSize = 5

type BingoBoard struct {
	nums []int
}

func NewBingoBoard(input string) BingoBoard {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanWords)

	nums := make([]int, 0)

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		ch(err)

		nums = append(nums, n)
	}
	return BingoBoard{nums}
}

func ReadBingoInput(input string) (boards []*BingoBoard) {
	ss := strings.Split(input, "\n\n")
	for _, s := range ss[1:] {
		b := NewBingoBoard(s)
		boards = append(boards, &b)
	}
	return
}

func ReadBingoNumOrder(input string) (order []int) {
	ss := strings.Split(input, "\n\n")
	nn := strings.Split(ss[0], ",")

	for _, n := range nn {
		num, err := strconv.Atoi(n)
		ch(err)
		order = append(order, num)
	}
	return
}

func (b *BingoBoard) MarkNumber(n int) {
	for i, num := range b.nums {
		if num == n {
			if num > 0 {
				b.nums[i] *= -1
			} else if num == 0 {
				b.nums[i] = -1
			}
		}
	}
}

func (b *BingoBoard) Wins() bool {
	for _, row := range b.Rows() {
		marked := 0
		for _, num := range row {
			if num < 0 {
				marked++
			}
		}
		if marked == BingoSize {
			return true
		}
	}
	for _, col := range b.Cols() {
		marked := 0
		for _, num := range col {
			if num < 0 {
				marked++
			}
		}
		if marked == BingoSize {
			return true
		}
	}
	return false
}

func (b *BingoBoard) Score(lastCalledNum int) int {
	var unmarked []int

	for _, num := range b.nums {
		if num > 0 {
			unmarked = append(unmarked, num)
		}
	}

	var sum int
	for _, n := range unmarked {
		sum += n
	}

	return sum * lastCalledNum
}

func (b *BingoBoard) Rows() (rows [][]int) {
	for i := 0; i < len(b.nums); i += BingoSize {
		end := i + BingoSize

		if end > len(b.nums) {
			end = len(b.nums)
		}

		rows = append(rows, b.nums[i:end])
	}
	return
}

func (b *BingoBoard) Cols() (cols [][]int) {
	r := b.Rows()

	for i := 0; i < BingoSize; i++ {
		col := make([]int, 0)
		for j := 0; j < BingoSize; j++ {
			col = append(col, r[j][i])
		}
		cols = append(cols, col)
	}

	return
}

func Day04() {
	b, err := ioutil.ReadFile("input/04.txt")
	ch(err)

	boards := ReadBingoInput(string(b))
	nums := ReadBingoNumOrder(string(b))

	var lastNum int
	var lastWinner *BingoBoard

	var allBoardsWon = func(boards []*BingoBoard) bool {
		for _, b := range boards {
			if !b.Wins() {
				return false
			}
		}
		return true
	}

	for _, n := range nums {
		// log.Println(n)
		for _, b := range boards {
			b.MarkNumber(n)
			// log.Println(b.Wins())
			if b.Wins() {
				if lastWinner == nil {
					fmt.Println("Part 1:", b.Score(n))
				}
				lastNum = n
				lastWinner = b
				if allBoardsWon(boards) {
					fmt.Println("Part 2:", lastWinner.Score(lastNum))
					os.Exit(0)
				}
			}
		}
	}

}
