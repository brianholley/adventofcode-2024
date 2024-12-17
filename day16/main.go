package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"math"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

type Cell struct {
	row int
	col int
}

const north = 1
const east = 2
const south = 3
const west = 4
const wall = "#"

func rotate(inDir int, outDir int) int {
	if inDir == outDir {
		return 0
	}
	if lib.Abs(inDir-outDir) == 2 {
		return 2000
	}
	return 1000 // 90 turn by default
}

func move(maze [][]string, paths [][]map[int]int, row int, col int, dir int) {
	// North
	if row > 0 && maze[row-1][col] != wall {
		dist := paths[row][col][dir] + 1 + rotate(dir, north)
		check, exists := paths[row-1][col][north]
		if !exists || dist < check {
			paths[row-1][col][north] = dist
			move(maze, paths, row-1, col, north)
		}
	}
	// East
	if col < len(paths[row])-1 && maze[row][col+1] != wall {
		dist := paths[row][col][dir] + 1 + rotate(dir, east)
		check, exists := paths[row][col+1][east]
		if !exists || dist < check {
			paths[row][col+1][east] = dist
			move(maze, paths, row, col+1, east)
		}
	}
	// South
	if row < len(paths)-1 && maze[row+1][col] != wall {
		dist := paths[row][col][dir] + 1 + rotate(dir, south)
		check, exists := paths[row+1][col][south]
		if !exists || dist < check {
			paths[row+1][col][south] = dist
			move(maze, paths, row+1, col, south)
		}
	}
	// West
	if col > 0 && maze[row][col-1] != wall {
		dist := paths[row][col][dir] + 1 + rotate(dir, west)
		check, exists := paths[row][col-1][west]
		if !exists || dist < check {
			paths[row][col-1][west] = dist
			move(maze, paths, row, col-1, west)
		}
	}
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	maze := [][]string{}
	start := Cell{}
	end := Cell{}
	for stdin.Scan() {
		line := stdin.Text()
		row := strings.Split(line, "")
		if lib.ArrayContains(row, "S") {
			start.row = len(maze)
			start.col = lib.ArrayIndexOf(row, "S")
		}
		if lib.ArrayContains(row, "E") {
			end.row = len(maze)
			end.col = lib.ArrayIndexOf(row, "E")
		}
		maze = append(maze, row)
	}

	paths := make([][]map[int]int, len(maze))
	for i := range paths {
		paths[i] = make([]map[int]int, len(maze[i]))
		for j := range paths[i] {
			paths[i][j] = map[int]int{}
		}
	}
	paths[start.row][start.col][east] = 0

	move(maze, paths, start.row, start.col, east)

	result = math.MaxInt
	for _, dist := range paths[end.row][end.col] {
		result = lib.Min(result, dist)
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
