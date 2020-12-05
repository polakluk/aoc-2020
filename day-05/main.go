package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// SeatPos holds information about a seat
type SeatPos struct {
	row int
	col int
}

// calcBisect calculates the value based on codes
func calcBisect(code string, lowerCmd rune) int {
	lower, middle, upper := 0, 0, int(math.Pow(2, float64(len(code))))-1

	for _, cmd := range code {
		middle = (upper + lower) / 2
		// offset to handle float numbers in middle point (127/2 = 63.5)
		offset := (upper+lower)%2 == 0

		if rune(cmd) == lowerCmd {
			if offset {
				middle--
			}
			upper = middle

		} else {
			if !offset {
				middle++
			}
			lower = middle
		}
	}

	return middle

}

// calcSeatInfo calculates seat information for given code
func calcSeatInfo(code string) SeatPos {
	return SeatPos{row: calcBisect(code[0:7], rune('F')), col: calcBisect(code[7:10], rune('L'))}
}

func calcSeatID(pos SeatPos) int {
	return pos.row*8 + pos.col
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// list of seat IDs
	var seatIDs []int

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		info := calcSeatInfo(line)
		ID := calcSeatID(info)

		seatIDs = append(seatIDs, ID)
	}

	sort.Ints(seatIDs)

	fmt.Printf("Highest seat ID - %d\n", seatIDs[len(seatIDs)-1])

	// find my seat
	for idx := 0; idx < len(seatIDs)-1; idx++ {
		if seatIDs[idx]+2 == seatIDs[idx+1] {
			fmt.Printf("My seat ID is %d\n", seatIDs[idx]+1)
		}
	}
}
