package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

const north = 1
const east = 2
const south = 3
const west = 4

type Fence struct {
	row       int
	col       int
	direction int
	length    int
}

func mapRegion(garden [][]string, counted [][]bool, row int, col int) (int, []Fence) {
	plant := garden[row][col]
	area := 1
	perimeter := []Fence{}
	counted[row][col] = true

	if row > 0 && garden[row-1][col] == plant {
		if !counted[row-1][col] {
			a, p := mapRegion(garden, counted, row-1, col)
			area += a
			perimeter = append(perimeter, p...)
		}
	} else {
		perimeter = append(perimeter, Fence{row: row, col: col, direction: north, length: 1})
	}

	if col < len(garden[row])-1 && garden[row][col+1] == plant {
		if !counted[row][col+1] {
			a, p := mapRegion(garden, counted, row, col+1)
			area += a
			perimeter = append(perimeter, p...)
		}
	} else {
		perimeter = append(perimeter, Fence{row: row, col: col, direction: east, length: 1})
	}

	if row < len(garden)-1 && garden[row+1][col] == plant {
		if !counted[row+1][col] {
			a, p := mapRegion(garden, counted, row+1, col)
			area += a
			perimeter = append(perimeter, p...)
		}
	} else {
		perimeter = append(perimeter, Fence{row: row, col: col, direction: south, length: 1})
	}

	if col > 0 && garden[row][col-1] == plant {
		if !counted[row][col-1] {
			a, p := mapRegion(garden, counted, row, col-1)
			area += a
			perimeter = append(perimeter, p...)
		}
	} else {
		perimeter = append(perimeter, Fence{row: row, col: col, direction: west, length: 1})
	}

	return area, perimeter
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	garden := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		garden = append(garden, strings.Split(line, ""))
	}

	counted := [][]bool{}
	for i := range garden {
		counted = append(counted, make([]bool, len(garden[i])))
	}

	for i := range garden {
		for j := range garden[i] {
			if !counted[i][j] {
				area, perimeter := mapRegion(garden, counted, i, j)
				// fmt.Println("Region", garden[i][j], "at", i, j, "has area", area, "and perimenter", perimeter)
				result += area * len(perimeter)
			}
		}
	}

	return fmt.Sprint(result)
}

func combineFences(fence []Fence) []Fence {
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(fence)-1; i++ {
			for j := i + 1; j < len(fence); j++ {
				if fence[i].direction == north || fence[i].direction == south {
					if fence[i].direction == fence[j].direction &&
						fence[i].row == fence[j].row &&
						(fence[i].col+fence[i].length == fence[j].col || fence[i].col == fence[j].col+fence[j].length) {

						fence[i].col = int(math.Min(float64(fence[i].col), float64(fence[j].col)))
						fence[i].length += fence[j].length
						fence = slices.Delete(fence, j, j+1)
						merged = true
					}
				} else {
					if fence[i].direction == fence[j].direction &&
						fence[i].col == fence[j].col &&
						(fence[i].row+fence[i].length == fence[j].row || fence[i].row == fence[j].row+fence[j].length) {

						fence[i].row = int(math.Min(float64(fence[i].row), float64(fence[j].row)))
						fence[i].length += fence[j].length
						fence = slices.Delete(fence, j, j+1)
						merged = true
					}
				}
			}
		}
	}
	return fence
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	garden := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		garden = append(garden, strings.Split(line, ""))
	}

	counted := [][]bool{}
	for i := range garden {
		counted = append(counted, make([]bool, len(garden[i])))
	}

	for i := range garden {
		for j := range garden[i] {
			if !counted[i][j] {
				area, perimeter := mapRegion(garden, counted, i, j)
				// fmt.Println("Region", garden[i][j], "at", i, j, "has area", area, "and perimenter", perimeter)
				perimeter = combineFences(perimeter)
				result += area * len(perimeter)
			}
		}
	}

	return fmt.Sprint(result)
}
