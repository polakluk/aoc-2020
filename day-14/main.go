package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func preprocessMask(mask string) string {
	for idx, val := range mask {
		if val != 'X' {
			return mask[idx:]
		}
	}
	panic("Error")
}

func applyMask(mask string, inp int) int64 {
	strVal := strconv.FormatInt(int64(inp), 2)

	finalVal := ""
	for idx := len(mask) - 1; idx >= 0; idx-- {
		digit := mask[idx]
		if digit == 'X' {
			pos := len(strVal) + idx - len(mask)
			if pos < 0 {
				digit = '0'
			} else {
				digit = strVal[pos]
			}
		}
		finalVal = string(digit) + finalVal
	}
	binVal, _ := strconv.ParseInt(finalVal, 2, 64)

	return binVal
}

func padding(inp string, expectedLen int) string {
	padZeros := ""
	for len(padZeros)+len(inp) < expectedLen {
		padZeros += "0"
	}
	return padZeros + inp
}

func calculatePositions(mask string, memory int) []string {
	positions := []string{}

	strVal := strconv.FormatInt(int64(memory), 2)

	maskedVal := ""
	numWildcards := 0
	wildcardsPos := []int{}
	for idx := len(mask) - 1; idx >= 0; idx-- {
		digit := mask[idx]
		if digit == '0' {
			pos := len(strVal) + idx - len(mask)
			if pos < 0 {
				digit = '0'
			} else {
				digit = strVal[pos]
			}
		} else {
			if digit == 'X' {
				numWildcards++
				wildcardsPos = append(wildcardsPos, idx)
			}
		}
		maskedVal = string(digit) + maskedVal
	}
	maxCombinations := 1 << numWildcards
	for val := 0; val < maxCombinations; val++ {
		actPosition := []rune(maskedVal)
		actFloatingVal := padding(strconv.FormatInt(int64(val), 2), numWildcards)

		for idx, pos := range wildcardsPos {
			actPosition[pos] = rune(actFloatingVal[idx])
		}

		positions = append(positions, padding(string(actPosition), 36))
	}

	return positions
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var actMask1, actMask2 string
	var memoryPosition int
	var val int
	memory := map[int]int64{}
	memory2 := map[string]int{}

	for scanner.Scan() {

		// read the line
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			actMask1 = preprocessMask(parts[1])
			actMask2 = parts[1]
		} else {
			fmt.Sscanf(line, "mem[%d] = %d", &memoryPosition, &val)

			// part 1
			memory[memoryPosition] = applyMask(actMask1, val)

			// part 2
			positions := calculatePositions(actMask2, memoryPosition)
			for _, pos := range positions {
				memory2[pos] = val
			}
		}
	}
	var totalSum int64 = 0
	for _, memVal := range memory {
		totalSum += memVal
	}

	fmt.Printf("Total sum v1 - %d\n", totalSum)

	var totalSum2 int64 = 0
	for _, memVal := range memory2 {
		totalSum2 += int64(memVal)
	}

	fmt.Printf("Total sum v2 - %d\n", totalSum2)
}
