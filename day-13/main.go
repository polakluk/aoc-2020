package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findCommonNaive(busIDs []int) int {
	t := busIDs[0]
	for {
		stopLooking := false
		for idx, val := range busIDs[1:] {
			if val == -1 {
				continue
			}
			if (t+idx+1)%val > 0 {
				stopLooking = true
				break
			}
		}

		if !stopLooking {
			break
		}
		t += busIDs[0]
	}

	return t
}

func findCommonNormal(busIDs []int) int {
	step := busIDs[0]
	actT := step
	actIdx := 1

	for ; actIdx < len(busIDs); actT += step {
		// skip "x"
		for busIDs[actIdx] == -1 {
			actIdx++
		}
		// check whether the current number is divisible by the current busId
		if (actT+actIdx)%busIDs[actIdx] == 0 {
			// if it is, then start skipping "actStep" + "actBusId" numbers and search for another match
			actIdx++
			if actIdx == len(busIDs) {
				// this was the last bus so return the value
				return actT
			}
			step *= busIDs[actIdx-1]
		}
	}
	// error state (should have been panic)
	return -1
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	arrivalTime, _ := strconv.Atoi(scanner.Text())

	// list of ids
	scanner.Scan()
	rawIds := scanner.Text()
	ids := strings.Split(rawIds, ",")

	waitTime := -1
	catchBusID := -1
	buses := []int{}
	for _, rawBusID := range ids {
		// skip malfunctioned buses
		if rawBusID == "x" {
			buses = append(buses, -1)
			continue
		}
		busID, _ := strconv.Atoi(rawBusID)
		buses = append(buses, busID)
		// calculate wait time for this bus
		actWaitTime := busID - (arrivalTime % busID)
		if catchBusID == -1 || waitTime > actWaitTime {
			// either this is the first bus we check or its wait time is lower than any wait time found before
			catchBusID = busID
			waitTime = actWaitTime
		}
	}

	fmt.Printf("Part 1 answer - %d\n", catchBusID*waitTime)
	fmt.Printf("Part 2 (normal) - %d\n", findCommonNormal(buses))
	fmt.Printf("Part 2 (naive) -  %d\n", findCommonNaive(buses))
}
