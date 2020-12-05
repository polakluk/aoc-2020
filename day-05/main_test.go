package main

import (
	"fmt"
	"testing"
)

type testCasePos struct {
	code     string
	expected SeatPos
}

func TestSeatPos(t *testing.T) {
	cases := []testCasePos{
		{"FBFBBFFRLR", SeatPos{row: 44, col: 5}},
		{"BFFFBBFRRR", SeatPos{row: 70, col: 7}},
		{"FFFBBBFRRR", SeatPos{row: 14, col: 7}},
		{"BBFFBBFRLL", SeatPos{row: 102, col: 4}},
	}
	for _, c := range cases {
		result := calcSeatInfo(c.code)
		if c.expected.row != result.row || c.expected.col != result.col {
			t.Errorf("Test case %s ==> Expected [%d, %d], got [%d, %d]", c.code, c.expected.col, c.expected.row, result.col, result.row)
		}
		fmt.Println(c)
	}
}

type testCaseBisect struct {
	code     string
	lowerCmd rune
	expected int
}

func TestBisect(t *testing.T) {
	cases := []testCaseBisect{
		{"FBFBBFF", 'F', 44},
		{"BFFFBBF", 'F', 70},
		{"FFFBBBF", 'F', 14},
		{"BBFFBBF", 'F', 102},

		{"RLR", 'L', 5},
		{"RRR", 'L', 7},
		{"RLL", 'L', 4},
	}
	for _, c := range cases {
		result := calcBisect(c.code, c.lowerCmd)
		if c.expected != result {
			t.Errorf("Test case %s ==> Expected %d], got %d", c.code, c.expected, result)
		}
		fmt.Println(c)
	}
}
