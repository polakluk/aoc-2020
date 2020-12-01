package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkPair(exp map[int]bool, actValue int, currentMulti int) int {
	// calculate hash of the desired key
	wantedKey := 2020 - actValue
	_, ok := exp[wantedKey]
	if ok {
		// if we found the key, return multiplication of the two numbers
		return wantedKey * actValue
	} else {
		// otherwise, return the known value for multiplication (in other words, do nothing)
		return currentMulti
	}
}

func checkThree(exp map[int]bool, actValue int, currentMulti int) int {
	// calculate hash of the desired key
	reminder := 2020 - actValue

	for firstVal := range exp {
		wantedKey := reminder - firstVal
		_, ok := exp[wantedKey]
		if ok {
			// if we found the key, return multiplication of the three numbers
			return actValue * firstVal * wantedKey
		}
	}

	// otherwise, return the known value for multiplication (in other words, do nothing)
	return currentMulti
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	expenses := make(map[int]bool)
	twoMulti, threeeMulti := 0, 0
	for scanner.Scan() {
		// load number from file
		val, _ := strconv.Atoi(scanner.Text())
		// check for mirror value to get sum 2020 and update pair multiplication accordingly
		twoMulti = checkPair(expenses, val, twoMulti)
		// check for another 2 values to get sum 2020 and update multiplication for three numbers  accordingly
		threeeMulti = checkThree(expenses, val, threeeMulti)

		// insert expense into the map
		expenses[val] = true
	}

	fmt.Printf("Pair Multiplication is %d \n", twoMulti)
	fmt.Printf("Three-numver multiplication is %d \n", threeeMulti)
}
