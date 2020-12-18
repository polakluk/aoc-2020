package main

import (
	"fmt"
)

func afterNTurnes(inp []int, maxTurns int) int {
	spoken := map[int][2]int{}

	for idx, val := range inp {
		spoken[val] = [2]int{idx + 1, -1}
	}
	previousNum := inp[len(inp)-1]

	for turn := len(spoken) + 1; turn <= maxTurns; turn++ {
		val, ok := spoken[previousNum]
		if ok {
			// this number was already mentioned
			nextVal := 0
			if val[1] != -1 {
				nextVal = val[0] - val[1]
			}

			valNextVal, okNextVal := spoken[nextVal]
			if okNextVal {
				// already found nextVal
				spoken[nextVal] = [2]int{turn, valNextVal[0]}
			} else {
				spoken[nextVal] = [2]int{turn, -1}
			}
			previousNum = nextVal
		} else {
			// this number was never mentioned
			val0, ok0 := spoken[0]
			if ok0 {
				// already found 0
				val0[1] = val0[0]
				val0[0] = turn
			} else {
				spoken[0] = [2]int{turn, -1}
			}
			previousNum = 0
		}
	}

	return previousNum
}

func main() {
	inp := []int{18, 8, 0, 5, 4, 1, 20}

	fmt.Printf("2020th spoken number is %d\n", afterNTurnes(inp, 2020))
	fmt.Printf("30000000th spoken number is %d\n", afterNTurnes(inp, 30000000))
}
