package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	for _, rawBusID := range ids {
		// skip malfunctioned buses
		if rawBusID == "x" {
			continue
		}
		busID, _ := strconv.Atoi(rawBusID)
		// calculate wait time for this bus
		actWaitTime := busID - (arrivalTime % busID)
		if catchBusID == -1 || waitTime > actWaitTime {
			// either this is the first bus we check or its wait time is lower than any wait time found before
			catchBusID = busID
			waitTime = actWaitTime
		}
	}

	fmt.Printf("Part 1 answer - %d\n", catchBusID*waitTime)
}
