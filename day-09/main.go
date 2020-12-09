package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	preamble := 25
	incorrectNumber := 0
	smallest := 0
	largest := 0
	inp := []int{}

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		inp = append(inp, num)
		if len(inp) > preamble {
			foundPair := false

			// check sums of combinations of last preamble numbers
			for outerIdx := len(inp) - preamble - 1; outerIdx < len(inp)-2 && !foundPair; outerIdx++ {
				for innerIdx := outerIdx + 1; innerIdx < len(inp)-1 && !foundPair; innerIdx++ {
					if inp[outerIdx]+inp[innerIdx] == num {
						foundPair = true
					}
				}
			}

			if !foundPair && incorrectNumber == 0 {
				// this is the first number which is not a sum of 2 numbers from last N numbers
				incorrectNumber = num

				// check sums of combinations of last preamble numbers
				foundSet := false
				for outerIdx := 0; outerIdx < len(inp)-2 && !foundSet; outerIdx++ {
					smallest = inp[outerIdx]
					largest = inp[outerIdx]
					actSum := inp[outerIdx]

					for innerIdx := outerIdx + 1; innerIdx < len(inp)-1 && !foundSet; innerIdx++ {
						actSum += inp[innerIdx]

						if inp[innerIdx] < smallest {
							smallest = inp[innerIdx]
						}
						if inp[innerIdx] > largest {
							largest = inp[innerIdx]
						}

						if actSum == num {
							// found the subset
							foundSet = true
						} else {
							if actSum > num {
								// sum is already larger than the number - abort
								break
							}
						}
					}
				}

			}
		}
	}

	fmt.Printf("The first number which validates the rules - %d\n", incorrectNumber)
	fmt.Printf("Weaknes number - %d\n", largest+smallest)

}
