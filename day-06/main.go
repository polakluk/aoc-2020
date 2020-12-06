package main

import (
	"bufio"
	"fmt"
	"os"
)

// evaluateGroup evaluates all answers within a group
func evaluateGroup(answers map[rune]int, groupSize int) (int, int) {
	allAnswered := 0
	for _, answered := range answers {
		if answered == groupSize {
			allAnswered++
		}
	}
	return len(answers), allAnswered
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	answers := make(map[rune]int)
	totalAnswers := 0
	allAnswered := 0

	groupSize := 0
	for scanner.Scan() {
		// read the line
		line := scanner.Text()

		if line == "" {
			// evaluate the group
			groupAnswers, groupAllAnswers := evaluateGroup(answers, groupSize)

			totalAnswers += groupAnswers
			allAnswered += groupAllAnswers

			// reset the map of answers
			answers = make(map[rune]int)
			groupSize = 0
		} else {
			groupSize++
			for _, c := range line {
				// insert each answer to the map
				_, ok := answers[rune(c)]
				if ok {
					answers[rune(c)]++
				} else {
					answers[rune(c)] = 1
				}
			}
		}
	}
	// evaluate the group
	groupAnswers, groupAllAnswers := evaluateGroup(answers, groupSize)

	totalAnswers += groupAnswers
	allAnswered += groupAllAnswers

	fmt.Printf("Total answers - %d\n", totalAnswers)
	fmt.Printf("Total all answered - %d\n", allAnswered)
}
