package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func checkAdaptersBrute(prevNum int, adapters []int, offset int) int {
	// if you're checking the last adapter, you always found a valid combination
	if offset == len(adapters) {
		return 1
	}
	keepLooking := true
	idx := offset
	foundCombinations := 0
	for keepLooking {
		if idx == len(adapters) || adapters[idx]-prevNum > 3 {
			keepLooking = false
		} else {
			foundCombinations += checkAdaptersBrute(adapters[idx], adapters, idx+1)

			idx++
		}
	}

	return foundCombinations
}

func checkAdaptersDynamic(adapters []int) int {
	// prepare the intermediate results
	results := []int{}
	for range adapters {
		results = append(results, -1)
	}
	results[0] = 1
	// calcualte the result using intermediate results
	return checkAdaptersDynamicWalk(adapters, results, len(adapters)-1)
}

func checkAdaptersDynamicWalk(adapters []int, results []int, idx int) int {
	if results[idx] == -1 {
		// if we dont know the answer for this position yet, calculate it from the partial results
		keepLooking := true
		combinatios := 0
		// check the intermediate results for all compatible adapters (within range of 3 from the current adapter)
		for actIdx := idx - 1; keepLooking && actIdx > -1; actIdx-- {
			if adapters[idx]-adapters[actIdx] > 3 {
				// no further combinations are possible
				keepLooking = false
			} else {
				// check the number of combinations for this adapter
				combinatios += checkAdaptersDynamicWalk(adapters, results, actIdx)
			}
		}
		results[idx] = combinatios
	}
	return results[idx]
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	diffs := map[int]int{1: 0, 2: 0, 3: 0}
	// add 0 at the beginning of list of adapters to make the code easier to read :)
	adapters := []int{0}

	prevNum := 0
	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)

	for _, num := range adapters {
		diffs[num-prevNum]++
		prevNum = num
	}
	// the final adapter
	diffs[3]++

	fmt.Printf("Part 1 - %d * %d = %d\n", diffs[1], diffs[3], diffs[1]*diffs[3])

	allAdaptersConnections := checkAdaptersDynamic(adapters)
	fmt.Printf("Part 2a - %d\n", allAdaptersConnections)
}
