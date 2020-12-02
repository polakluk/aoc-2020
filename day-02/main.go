package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkOldPassword(inp string, character byte, lowerBoundary int, upperBoundary int) bool {
	occurence := 0
	for idx := 0; idx < len(inp); idx++ {
		if inp[idx] == character {
			occurence++
		}
	}
	if lowerBoundary <= occurence && occurence <= upperBoundary {
		return true
	} else {
		return false
	}
}

func checkNewPassword(inp string, character byte, firstPos int, secondPos int) bool {
	isFirstPos := inp[firstPos-1] == character
	isSecondPos := inp[secondPos-1] == character

	return (isFirstPos || isSecondPos) && !(isFirstPos && isSecondPos)

}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var inp string
	var character rune
	var lowerBoundary, upperBoundary int
	cntOldPass := 0
	cntNewPass := 0

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		// detect parameters in the line
		fmt.Sscanf(line, "%d-%d %c:%s", &lowerBoundary, &upperBoundary, &character, &inp)
		// check the old password format
		if checkOldPassword(inp, byte(character), lowerBoundary, upperBoundary) {
			cntOldPass++
		}
		// check the new password format
		if checkNewPassword(inp, byte(character), lowerBoundary, upperBoundary) {
			cntNewPass++
		}
	}
	fmt.Printf("Valid old passwords - %d\n", cntOldPass)
	fmt.Printf("Valid new passwords - %d\n", cntNewPass)

}
