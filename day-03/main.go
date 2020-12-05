package main

import (
	"bufio"
	"fmt"
	"os"
)

// SlidingScenarioInput is definition of sliding scenario
type SlidingScenarioInput struct {
	right int
	down  int
}

// SlidingScenarioResult is result for a sliding scenario
type SlidingScenarioResult struct {
	currentCol int
	foundTrees int
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// starting line number
	lineNum := 0
	// sliding scenarios
	scenarios := [5]SlidingScenarioInput{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	// sliding scenarios results
	results := [5]SlidingScenarioResult{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	for scanner.Scan() {
		// read the line
		line := scanner.Text()

		// skip the checks for the first line
		if lineNum > 0 {
			// check all scenarios
			for idx := 0; idx < 5; idx++ {
				// we perform the movement and the check every "down"-ed line
				shouldMoveAndCheck := lineNum%scenarios[idx].down == 0

				if shouldMoveAndCheck {
					// move to the right
					results[idx].currentCol = (results[idx].currentCol + scenarios[idx].right) % len(line)

					if rune(line[results[idx].currentCol]) == '#' {
						// we performed the check and it turns out, we hit a tree :( (we skip the check on the first line)
						results[idx].foundTrees++
					}
				}
			}
		}

		lineNum++
	}

	multiplyTrees := 1
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("Found trees for scenario %d - %d\n", idx, results[idx].foundTrees)
		multiplyTrees *= results[idx].foundTrees
	}

	fmt.Printf("\nOriginal scenario trees - %d\n", results[1].foundTrees)
	fmt.Printf("Multiplied number of trees - %d\n", multiplyTrees)
}
