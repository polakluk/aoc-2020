package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// moveAction returns offset of the movement and the new facing direction
func moveAction(act rune, facing rune, length float64) (float64, float64, rune) {
	switch act {
	case 'N':
		{
			return 0, -length, facing
		}
	case 'S':
		{
			return 0, length, facing
		}
	case 'E':
		{
			return -length, 0, facing
		}
	case 'W':
		{
			return length, 0, facing
		}
	case 'L':
		{
			angle := length / 90
			actPosition := strings.Index("NWSE", string(facing))
			newAngle := (actPosition + int(angle)) % 4
			return 0.0, 0.0, rune("NWSE"[newAngle])
		}
	case 'R':
		{
			angle := length / 90
			actPosition := strings.Index("NESW", string(facing))
			newAngle := (actPosition + int(angle)) % 4
			return 0.0, 0.0, rune("NESW"[newAngle])
		}
	case 'F':
		{
			// only move along the facing direction
			return moveAction(facing, facing, length)
		}
	}
	panic("This should never happen :D")
}

// moveWaypoint returns offset of the movement and the new waypoint
func moveWaypoint(act rune, waypoint [2]float64, length float64) ([2]float64, [2]float64) {
	switch act {
	case 'N':
		{
			return [2]float64{0.0, 0.0}, [2]float64{waypoint[0] - length, waypoint[1]}
		}
	case 'S':
		{
			return [2]float64{0.0, 0.0}, [2]float64{waypoint[0] + length, waypoint[1]}
		}
	case 'E':
		{
			return [2]float64{0.0, 0.0}, [2]float64{waypoint[0], waypoint[1] - length}
		}
	case 'W':
		{
			return [2]float64{0.0, 0.0}, [2]float64{waypoint[0], waypoint[1] + length}
		}
	case 'L':
		{
			angle := length / 90
			newWaypoint := [2]float64{waypoint[0], waypoint[1]}
			for idx := 0.0; idx < angle; idx++ {
				tmp := newWaypoint[0]
				newWaypoint[0] = newWaypoint[1]
				newWaypoint[1] = -tmp
			}
			return [2]float64{0.0, 0.0}, newWaypoint
		}
	case 'R':
		{
			// turn this rotation to counter-clockwise
			angle := 4.0 - length/90
			newWaypoint := [2]float64{waypoint[0], waypoint[1]}
			for idx := 0.0; idx < angle; idx++ {
				tmp := newWaypoint[0]
				newWaypoint[0] = newWaypoint[1]
				newWaypoint[1] = -tmp
			}
			return [2]float64{0.0, 0.0}, newWaypoint

		}
	case 'F':
		{
			return [2]float64{waypoint[0] * length, waypoint[1] * length}, waypoint
		}
	}
	panic("This should never happen :D")
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	// start facing East
	currentDirection := 'E'
	var actionCode rune
	diffRow, diffCol, length := 0.0, 0.0, 0.0

	// starting position (row, col)
	position := [2][2]float64{{0.0, 0.0}, {0.0, 0.0}}
	waypoint := [2]float64{-1.0, -10.0}
	var positionMovement [2]float64

	for scanner.Scan() {

		// read the line
		line := scanner.Text()

		// detect parameters in the line
		fmt.Sscanf(line, "%c%f", &actionCode, &length)
		// part 1
		diffRow, diffCol, currentDirection = moveAction(actionCode, currentDirection, length)
		position[0][0] += diffRow
		position[0][1] += diffCol

		// fmt.Printf("Code = %c; Length = %.0f; Position = [%.0f, %.0f]; NewDir = %c\n", actionCode, length, position[0][0], position[0][1], currentDirection)

		// part 2
		positionMovement, waypoint = moveWaypoint(actionCode, waypoint, length)
		position[1][0] += positionMovement[0]
		position[1][1] += positionMovement[1]
		// fmt.Printf(
		// 	"%c[%.0f]; Position = [%.0f, %.0f]; Waypoint = [%.0f, %.0f]\n",
		// 	actionCode,
		// 	length,
		// 	position[1][0],
		// 	position[1][1],
		// 	waypoint[0],
		// 	waypoint[1],
		// )
	}

	fmt.Printf("Total distance is %.0f (part 1)\n", math.Abs(position[0][0])+math.Abs(position[0][1]))
	fmt.Printf("Total distance is %.0f (part 2)\n", math.Abs(position[1][0])+math.Abs(position[1][1]))
}
