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
	direction int
}

const north = 1
const east = 2
const south = 3
const west = 4

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

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
