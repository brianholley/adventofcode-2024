package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

type Guard struct {
	row       int
	col       int
	direction int
}

type cell struct {
	row int
	col int
}

const obstacle = -1
const north = 1
const east = 2
const south = 4
const west = 8

func readMap(stdin *bufio.Scanner) ([][]int, Guard) {
	areaMap := [][]int{}
	guard := Guard{}
	for i := 0; stdin.Scan(); i++ {
		line := stdin.Text()
		areaMap = append(areaMap, make([]int, len(line)))

		for j := range line {
			if line[j] == '#' {
				areaMap[i][j] = obstacle
			}
		}

		if strings.Contains(line, "^") {
			guard.row = i
			guard.col = strings.Index(line, "^")
			guard.direction = north
		}
	}
	return areaMap, guard
}

func runGuardCheckForLoops(guard Guard, areaMap [][]int) bool {
	mapWidth := len(areaMap)
	mapHeight := len(areaMap[0])

	for {
		if guard.row < 0 || guard.row >= mapHeight || guard.col < 0 || guard.col >= mapWidth {
			return false
		}
		if areaMap[guard.row][guard.col]&guard.direction == guard.direction {
			return true
		}
		areaMap[guard.row][guard.col] |= guard.direction
		if guard.direction == north {
			if guard.row > 0 && areaMap[guard.row-1][guard.col] == obstacle {
				guard.direction = east
			} else {
				guard.row--
			}
		} else if guard.direction == east {
			if guard.col < mapWidth-1 && areaMap[guard.row][guard.col+1] == obstacle {
				guard.direction = south
			} else {
				guard.col++
			}
		} else if guard.direction == south {
			if guard.row < mapHeight-1 && areaMap[guard.row+1][guard.col] == obstacle {
				guard.direction = west
			} else {
				guard.row++
			}
		} else if guard.direction == west {
			if guard.col > 0 && areaMap[guard.row][guard.col-1] == obstacle {
				guard.direction = north
			} else {
				guard.col--
			}
		}
	}
}

func part1(stdin *bufio.Scanner) string {
	result := 0

	areaMap, guard := readMap(stdin)
	runGuardCheckForLoops(guard, areaMap)

	for i := range areaMap {
		for j := range areaMap[i] {
			if areaMap[i][j] > 0 {
				result++
			}
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	originalMap, originalGuard := readMap(stdin)

	firstPassMap := lib.Array2dCopy(originalMap)
	firstPassGuard := originalGuard
	runGuardCheckForLoops(firstPassGuard, firstPassMap)

	objects := []cell{}

	for i := range originalMap {
		for j := range originalMap[i] {

			if firstPassMap[i][j] > 0 && (i != originalGuard.row || j != originalGuard.col) {
				areaMap := lib.Array2dCopy(originalMap)
				guard := originalGuard

				areaMap[i][j] = obstacle
				if runGuardCheckForLoops(guard, areaMap) {
					objects = append(objects, cell{row: i, col: j})
				}
			}
		}
	}

	// for _, o := range objects {
	// 	fmt.Println(o)
	// }
	result = len(objects)

	return fmt.Sprint(result)
}
