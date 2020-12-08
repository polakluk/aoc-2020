package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findUsedColors(node string, matrix map[string]map[string]int, colors map[string]int) {
	edges, ok := matrix[node]

	if ok {
		for nextNode := range edges {
			_, ok := colors[nextNode]
			if !ok {
				colors[nextNode] = 1
				findUsedColors(nextNode, matrix, colors)
			}
		}
	}
}

func calcCombinations(node string, matrix map[string]map[string]int) int {
	edges, ok := matrix[node]

	numBags := 0
	if ok {
		for nextNode, cntBags := range edges {
			numBags += cntBags

			numBags += cntBags * calcCombinations(nextNode, matrix)
		}
	}
	return numBags
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var adjecancyMatrix = map[string]map[string]int{}
	var invAdjecancyMatrix = map[string]map[string]int{}

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		parts := strings.Split(line, " bags contain ")

		// parse the other bags
		otherBags := strings.Split(parts[1], ", ")

		invEdges, ok := invAdjecancyMatrix[parts[0]]
		if !ok {
			// initialize the nested map
			invEdges = map[string]int{}
			invAdjecancyMatrix[parts[0]] = invEdges
		}

		if len(otherBags) > 1 || (len(otherBags) == 1 && otherBags[0] != "no other bags.") {
			for _, bagWithNumber := range otherBags {
				bagParts := strings.Split(bagWithNumber, " ")
				bag := strings.Join(bagParts[1:len(bagParts)-1], " ")

				// "source" bag
				edges, ok := adjecancyMatrix[bag]
				if !ok {
					// initialize the nested map
					edges = map[string]int{}
					adjecancyMatrix[bag] = edges
				}

				edges[parts[0]] = 1
				invEdges[bag], _ = strconv.Atoi(bagParts[0])
			}
		}
	}

	usedColors := make(map[string]int)
	findUsedColors("shiny gold", adjecancyMatrix, usedColors)
	combinations := calcCombinations("shiny gold", invAdjecancyMatrix)

	fmt.Printf("Number of used colors - %d\n", len(usedColors))
	fmt.Printf("Number of combinations - %d\n", combinations)
}
