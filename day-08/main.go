package main

import (
	"bufio"
	"fmt"
	"os"
)

type cmd struct {
	name  string
	param int
}

func runCode(cmds []cmd) (int, bool) {
	// keep track of visited commands
	usedCmds := make(map[int]bool)
	// keep executing th code?
	keepGoing := true
	// my act position
	actCmdIdx := 0
	// current value in accumulator
	acc := 0

	for keepGoing {
		if actCmdIdx == len(cmds) {
			keepGoing = false
		} else {
			_, ok := usedCmds[actCmdIdx]
			if ok {
				return acc, false
			} else {
				usedCmds[actCmdIdx] = true
				switch choice := cmds[actCmdIdx].name; choice {
				case "nop":
					actCmdIdx++
				case "jmp":
					actCmdIdx += cmds[actCmdIdx].param
				case "acc":
					acc += cmds[actCmdIdx].param
					actCmdIdx++
				}
			}
		}
	}

	return acc, true
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	cmds := []cmd{}
	var cmdName string
	var operator rune
	var param int
	var offset int

	for scanner.Scan() {
		// read the line
		line := scanner.Text()

		// detect parameters in the line
		fmt.Sscanf(line, "%s %c%d", &cmdName, &operator, &param)

		// keep track of the commands
		offset = param
		if operator == '-' {
			offset = -offset
		}
		actCmd := cmd{cmdName, offset}
		cmds = append(cmds, actCmd)
	}

	// run without change
	acc, _ := runCode(cmds)

	fmt.Printf("Accumulator value for the original code - %d\n", acc)

	// try to fix the code
	for idx := 0; idx < len(cmds); idx++ {
		originalCmdName := cmds[idx].name
		if originalCmdName != "acc" {
			newCmdName := "nop"
			if originalCmdName == "nop" {
				newCmdName = "jmp"
			}

			// change the operation
			cmds[idx].name = newCmdName

			// try to run the code
			acc, finished := runCode(cmds)

			if finished {
				fmt.Printf("Accumulator value after fixing code - %d (fixed instruction %d)\n", acc, idx)
				break
			}

			// revert the change
			cmds[idx].name = originalCmdName
		}
	}
}
