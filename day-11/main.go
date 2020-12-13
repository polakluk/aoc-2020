package main

import (
	"bufio"
	"fmt"
	"os"
)

// InitialSetup holds information about setup of the runner
type InitialSetup struct {
	mapIdx          int
	onlyOne         bool
	occupationLimit int
}

// ocuppiedAdjacent calculates number of adjacent occupied seats
func ocuppiedAdjacent(actMap map[int]map[int]rune, row int, col int, onlyOne bool) int {
	movements := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}
	occupied := 0

	for _, move := range movements {
		keepLooking := true

		for dist := 1; keepLooking; dist++ {
			val, ok := actMap[row+move[0]*dist][col+move[1]*dist]
			if ok {
				switch val {
				case '#':
					{
						occupied++
						keepLooking = false
					}
				case 'L':
					{
						keepLooking = false
					}
				}
			}

			// check only really adjacent seats
			if onlyOne || !ok {
				keepLooking = false
			}
		}
	}

	return occupied
}

// displayMap displays the map
func displayMap(actMap map[int]map[int]rune) {
	for actRow := 0; actRow < len(actMap); actRow++ {
		colMap := actMap[actRow]

		for actCol := 0; actCol < len(colMap); actCol++ {
			val := colMap[actCol]
			fmt.Printf("%c", val)
		}

		fmt.Print("\n")
	}
}

// countOccupiedSeats counts number of occupied seats
func countOccupiedSeats(actMap map[int]map[int]rune) int {
	seats := 0
	for actRow := 0; actRow < len(actMap); actRow++ {
		colMap := actMap[actRow]

		for actCol := 0; actCol < len(colMap); actCol++ {
			val := colMap[actCol]
			if val == '#' {
				seats++
			}
		}
	}
	return seats
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// just store 2 copies of the map for each part of the solution
	maps := [2][2]map[int]map[int]rune{{{}, {}}, {{}, {}}}

	row := 0
	for scanner.Scan() {
		// for part 1
		maps[0][0][row] = make(map[int]rune)
		maps[0][1][row] = make(map[int]rune)
		// for part 2
		maps[1][0][row] = make(map[int]rune)
		maps[1][1][row] = make(map[int]rune)

		// read the line
		line := scanner.Text()

		for idx, c := range line {
			maps[0][0][row][idx] = c
			maps[1][0][row][idx] = c
		}
		row++
	}

	displayMap(maps[0][0])
	fmt.Println("--------- SOLVING NOW ----------")
	parts := [2]InitialSetup{{0, true, 3}, {1, false, 4}}

	for idx, setup := range parts {
		fmt.Printf("Part %d:\n", idx+1)
		iterations := 0
		activeMap := 0
		nextMap := 1
		changes := -1

		for changes != 0 {
			changes = 0
			for actRow := 0; actRow < len(maps[setup.mapIdx][activeMap]); actRow++ {
				colMap := maps[setup.mapIdx][activeMap][actRow]
				for actCol := 0; actCol < len(colMap); actCol++ {

					switch val := colMap[actCol]; val {
					case '.':
						{
							maps[setup.mapIdx][nextMap][actRow][actCol] = '.'
						}
					case '#':
						{
							occupied := ocuppiedAdjacent(maps[setup.mapIdx][activeMap], actRow, actCol, setup.onlyOne)
							if occupied > setup.occupationLimit {
								maps[setup.mapIdx][nextMap][actRow][actCol] = 'L'
								changes++
							} else {
								maps[setup.mapIdx][nextMap][actRow][actCol] = '#'
							}
						}
					case 'L':
						{
							occupied := ocuppiedAdjacent(maps[setup.mapIdx][activeMap], actRow, actCol, setup.onlyOne)
							if occupied == 0 {
								maps[setup.mapIdx][nextMap][actRow][actCol] = '#'
								changes++
							} else {
								maps[setup.mapIdx][nextMap][actRow][actCol] = 'L'
							}

						}
					}
				}
			}

			// swap maps
			activeMap = 1 - activeMap
			nextMap = 1 - nextMap
			iterations++
		}

		seats := countOccupiedSeats(maps[setup.mapIdx][nextMap])
		fmt.Printf("After iterations - %d - we see %d seats\n", iterations, seats)
		displayMap(maps[setup.mapIdx][nextMap])

		fmt.Print("\n")
	}
}
