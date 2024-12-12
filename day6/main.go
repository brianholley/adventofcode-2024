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

type guardPos struct {
	row       int
	col       int
	direction string
}

type cell struct {
	row int
	col int
}

const north = "N"
const east = "E"
const south = "S"
const west = "W"

func part1(stdin *bufio.Scanner) string {
	result := 0

	areaMap := [][]string{}
	guard := guardPos{}
	for i := 0; stdin.Scan(); i++ {
		line := stdin.Text()
		areaMap = append(areaMap, strings.Split(line, ""))

		if strings.Contains(line, "^") {
			guard.row = i
			guard.col = strings.Index(line, "^")
			guard.direction = north
		}
	}

	mapWidth := len(areaMap)
	mapHeight := len(areaMap[0])

	for {
		if guard.row < 0 || guard.row >= mapHeight || guard.col < 0 || guard.col >= mapWidth {
			break
		}
		areaMap[guard.row][guard.col] = "x"
		if guard.direction == north {
			if guard.row > 0 && areaMap[guard.row-1][guard.col] == "#" {
				guard.direction = east
			} else {
				guard.row--
			}
		} else if guard.direction == east {
			if guard.col < mapWidth-1 && areaMap[guard.row][guard.col+1] == "#" {
				guard.direction = south
			} else {
				guard.col++
			}
		} else if guard.direction == south {
			if guard.row < mapHeight-1 && areaMap[guard.row+1][guard.col] == "#" {
				guard.direction = west
			} else {
				guard.row++
			}
		} else if guard.direction == west {
			if guard.col > 0 && areaMap[guard.row][guard.col-1] == "#" {
				guard.direction = north
			} else {
				guard.col--
			}
		}
	}

	for i := range areaMap {
		for j := range areaMap[i] {
			if areaMap[i][j] == "x" {
				result++
			}
		}
	}
	return fmt.Sprint(result)
}

// Brute force solution :(
func encountersLoop(areaMap [][]string, guard guardPos) bool {
	mapWidth := len(areaMap)
	mapHeight := len(areaMap[0])

	for {
		if guard.row < 0 || guard.row >= mapHeight || guard.col < 0 || guard.col >= mapWidth {
			return false
		}
		if strings.Contains(areaMap[guard.row][guard.col], guard.direction) {
			return true
		}
		areaMap[guard.row][guard.col] = areaMap[guard.row][guard.col] + guard.direction
		if guard.direction == north {
			if guard.row > 0 && areaMap[guard.row-1][guard.col] == "#" {
				guard.direction = east
			} else {
				guard.row--
			}
		} else if guard.direction == east {
			if guard.col < mapWidth-1 && areaMap[guard.row][guard.col+1] == "#" {
				guard.direction = south
			} else {
				guard.col++
			}
		} else if guard.direction == south {
			if guard.row < mapHeight-1 && areaMap[guard.row+1][guard.col] == "#" {
				guard.direction = west
			} else {
				guard.row++
			}
		} else if guard.direction == west {
			if guard.col > 0 && areaMap[guard.row][guard.col-1] == "#" {
				guard.direction = north
			} else {
				guard.col--
			}
		}
	}
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	originalMap := [][]string{}
	originalGuard := guardPos{}
	for i := 0; stdin.Scan(); i++ {
		line := stdin.Text()
		originalMap = append(originalMap, strings.Split(line, ""))

		if strings.Contains(line, "^") {
			originalGuard.row = i
			originalGuard.col = strings.Index(line, "^")
			originalGuard.direction = north
		}
	}

	objects := []cell{}

	for i := range originalMap {
		for j := range originalMap[i] {

			if originalMap[i][j] == "." {
				areaMap := lib.Array2dCopy(originalMap)
				guard := originalGuard

				areaMap[i][j] = "#"
				if encountersLoop(areaMap, guard) {
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
