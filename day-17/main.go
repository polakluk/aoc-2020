package main

import (
	"bufio"
	"fmt"
	"os"
)

func generateKey3D(row int, col int, depth int) string {
	return fmt.Sprintf("%d,%d,%d", col, row, depth)
}

func getCoordinate3D(coordinate string) (int, int, int) {
	var actCol, actRow, actDepth int
	fmt.Sscanf(coordinate, "%d,%d,%d", &actCol, &actRow, &actDepth)

	return actRow, actCol, actDepth
}

func generateKey4D(row int, col int, depth int, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", col, row, depth, w)
}

func getCoordinate4D(coordinate string) (int, int, int, int) {
	var actCol, actRow, actDepth, actW int
	fmt.Sscanf(coordinate, "%d,%d,%d,%d", &actCol, &actRow, &actDepth, &actW)

	return actRow, actCol, actDepth, actW
}

func neigborActiveCubes3D(cubeMap map[string]rune, row int, col, depth int) int {
	active := 0

	for actCol := -1; actCol < 2; actCol++ {
		for actRow := -1; actRow < 2; actRow++ {
			for actDepth := -1; actDepth < 2; actDepth++ {
				if actCol == 0 && actDepth == 0 && actRow == 0 {
					continue
				}
				coordinate := generateKey3D(actRow+row, actCol+col, actDepth+depth)
				val, ok := cubeMap[coordinate]
				if ok && val == '#' {
					active++
				}
			}
		}
	}
	return active
}

func addMissingNeighbors3D(cubeMap map[string]rune, row int, col, depth int) {
	for actCol := -1; actCol < 2; actCol++ {
		for actRow := -1; actRow < 2; actRow++ {
			for actDepth := -1; actDepth < 2; actDepth++ {
				if actCol == 0 && actDepth == 0 && actRow == 0 {
					continue
				}
				coordinate := generateKey3D(actRow+row, actCol+col, actDepth+depth)
				_, ok := cubeMap[coordinate]
				if !ok {
					cubeMap[coordinate] = '.'
				}
			}
		}
	}
}

func neigborActiveCubes4D(cubeMap map[string]rune, row int, col, depth int, w int) int {
	active := 0

	for actCol := -1; actCol < 2; actCol++ {
		for actRow := -1; actRow < 2; actRow++ {
			for actDepth := -1; actDepth < 2; actDepth++ {
				for actW := -1; actW < 2; actW++ {
					if actCol == 0 && actDepth == 0 && actRow == 0 && actW == 0 {
						continue
					}
					coordinate := generateKey4D(actRow+row, actCol+col, actDepth+depth, actW+w)
					val, ok := cubeMap[coordinate]
					if ok && val == '#' {
						active++
					}
				}
			}
		}
	}
	return active
}

func addMissingNeighbors4D(cubeMap map[string]rune, row int, col, depth int, w int) {
	for actCol := -1; actCol < 2; actCol++ {
		for actRow := -1; actRow < 2; actRow++ {
			for actDepth := -1; actDepth < 2; actDepth++ {
				for actW := -1; actW < 2; actW++ {
					if actCol == 0 && actDepth == 0 && actRow == 0 && actW == 0 {
						continue
					}
					coordinate := generateKey4D(actRow+row, actCol+col, actDepth+depth, actW+w)
					_, ok := cubeMap[coordinate]
					if !ok {
						cubeMap[coordinate] = '.'
					}
				}
			}
		}
	}
}

func cntActiveCubes(cubeMap map[string]rune) int {
	active := 0
	for _, val := range cubeMap {
		if val == '#' {
			active++
		}
	}
	return active
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	cubeMaps := [2]map[string]rune{{}, {}}
	cubeMaps4D := [2]map[string]rune{{}, {}}

	row := 0
	activeMap := 0
	nextMap := 1
	for scanner.Scan() {

		// read the line
		line := scanner.Text()
		for col, val := range line {
			coordinate := generateKey3D(row, col, 0)
			coordinate4D := generateKey4D(row, col, 0, 0)
			cubeMaps[activeMap][coordinate] = val
			cubeMaps4D[activeMap][coordinate4D] = val
			addMissingNeighbors3D(cubeMaps[activeMap], row, col, 0)
			addMissingNeighbors4D(cubeMaps4D[activeMap], row, col, 0, 0)
		}
		row++
	}

	for iteration := 1; iteration < 7; iteration++ {
		// part 1
		for key, val := range cubeMaps[activeMap] {
			actRow, actCol, actDepth := getCoordinate3D(key)
			neighbors := neigborActiveCubes3D(cubeMaps[activeMap], actRow, actCol, actDepth)
			if val == '#' {
				if neighbors == 2 || neighbors == 3 {
					cubeMaps[nextMap][key] = '#'
				} else {
					cubeMaps[nextMap][key] = '.'
				}
			} else {
				if neighbors == 3 {
					cubeMaps[nextMap][key] = '#'
				} else {
					cubeMaps[nextMap][key] = '.'
				}

			}
			addMissingNeighbors3D(cubeMaps[nextMap], actRow, actCol, actDepth)
		}

		// part 2
		for key, val := range cubeMaps4D[activeMap] {
			actRow, actCol, actDepth, actW := getCoordinate4D(key)
			neighbors := neigborActiveCubes4D(cubeMaps4D[activeMap], actRow, actCol, actDepth, actW)
			if val == '#' {
				if neighbors == 2 || neighbors == 3 {
					cubeMaps4D[nextMap][key] = '#'
				} else {
					cubeMaps4D[nextMap][key] = '.'
				}
			} else {
				if neighbors == 3 {
					cubeMaps4D[nextMap][key] = '#'
				} else {
					cubeMaps4D[nextMap][key] = '.'
				}

			}
			addMissingNeighbors4D(cubeMaps4D[nextMap], actRow, actCol, actDepth, actW)
		}

		activeMap = 1 - activeMap
		nextMap = 1 - nextMap
	}

	fmt.Printf("%d\n", cntActiveCubes(cubeMaps[activeMap]))
	fmt.Printf("%d\n", cntActiveCubes(cubeMaps4D[activeMap]))
}
