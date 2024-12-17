package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strings"
)

const wall = 1
const box = 2
const boxL = 3
const boxR = 4

type Robot struct {
	row, col int
}

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	result := 0

	warehouse := [][]int{}
	robot := Robot{}
	moves := []string{}
	for stdin.Scan() {
		line := stdin.Text()
		if strings.HasPrefix(line, "#") {
			row := make([]int, len(line))
			for i, s := range strings.Split(line, "") {
				if s == "#" {
					row[i] = wall
				} else if s == "O" {
					row[i] = box
				} else if s == "@" {
					robot.row = len(warehouse)
					robot.col = i
				}
			}
			warehouse = append(warehouse, row)
		} else {
			moves = append(moves, strings.Split(line, "")...)
		}
	}

	for _, move := range moves {
		switch move {
		case "^":
			if robot.row > 0 {
				if warehouse[robot.row-1][robot.col] == wall {
					continue
				}
				if warehouse[robot.row-1][robot.col] == box {
					i := robot.row - 1
					for ; i >= 0 && warehouse[i][robot.col] == box; i-- {
					}
					if warehouse[i][robot.col] == wall {
						continue
					}
					warehouse[i][robot.col] = box
					warehouse[robot.row-1][robot.col] = 0
				}
				robot.row--
			}
		case "v":
			if robot.row < len(warehouse)-1 {
				if warehouse[robot.row+1][robot.col] == wall {
					continue
				}
				if warehouse[robot.row+1][robot.col] == box {
					i := robot.row + 1
					for ; i < len(warehouse) && warehouse[i][robot.col] == box; i++ {
					}
					if warehouse[i][robot.col] == wall {
						continue
					}
					warehouse[i][robot.col] = box
					warehouse[robot.row+1][robot.col] = 0
				}
				robot.row++
			}
		case "<":
			if robot.col > 0 {
				if warehouse[robot.row][robot.col-1] == wall {
					continue
				}
				if warehouse[robot.row][robot.col-1] == box {
					j := robot.col - 1
					for ; j >= 0 && warehouse[robot.row][j] == box; j-- {
					}
					if warehouse[robot.row][j] == wall {
						continue
					}
					warehouse[robot.row][j] = box
					warehouse[robot.row][robot.col-1] = 0
				}
				robot.col--
			}
		case ">":
			if robot.col < len(warehouse[robot.row])-1 {
				if warehouse[robot.row][robot.col+1] == wall {
					continue
				}
				if warehouse[robot.row][robot.col+1] == box {
					j := robot.col + 1
					for ; j < len(warehouse[robot.row]) && warehouse[robot.row][j] == box; j++ {
					}
					if warehouse[robot.row][j] == wall {
						continue
					}
					warehouse[robot.row][j] = box
					warehouse[robot.row][robot.col+1] = 0
				}
				robot.col++
			}
		}
	}

	for i := range warehouse {
		for j := range warehouse[i] {
			if warehouse[i][j] == box {
				result += 100*i + j
			}
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	warehouse := [][]int{}
	robot := Robot{}
	moves := []string{}
	for stdin.Scan() {
		line := stdin.Text()
		if strings.HasPrefix(line, "#") {
			row := make([]int, len(line)*2)
			for j, s := range strings.Split(line, "") {
				if s == "#" {
					row[j*2] = wall
					row[j*2+1] = wall
				} else if s == "O" {
					row[j*2] = boxL
					row[j*2+1] = boxR
				} else if s == "@" {
					robot.row = len(warehouse)
					robot.col = j * 2
				}
			}
			warehouse = append(warehouse, row)
		} else {
			moves = append(moves, strings.Split(line, "")...)
		}
	}

	// fmt.Println("Start:")
	// printWarehouse(warehouse, robot)

	for _, move := range moves {
		robot = moveWithWideBoxes(move, robot, warehouse)
		// fmt.Println("After", move, ":")
		// printWarehouse(warehouse, robot)
		// fmt.Println()
	}

	// fmt.Println("End:")
	// printWarehouse(warehouse, robot)

	for i := range warehouse {
		for j := range warehouse[i] {
			if warehouse[i][j] == boxL {
				result += 100*i + j
			}
		}
	}

	return fmt.Sprint(result)
}

func canMoveBox(warehouse [][]int, row int, col int, direction string) bool {
	switch direction {
	case "^":
		if row > 0 {
			canMove := true
			if warehouse[row-1][col] == wall || warehouse[row-1][col+1] == wall {
				return false
			}
			if warehouse[row-1][col] == boxL || warehouse[row-1][col] == boxR {
				canMove = canMove && canMoveBox(warehouse, row-1, col-(warehouse[row-1][col]-boxL), direction)
			}
			if warehouse[row-1][col+1] == boxL {
				canMove = canMove && canMoveBox(warehouse, row-1, col+1, direction)
			}
			return canMove
		}
	case "v":
		if row < len(warehouse)-1 {
			canMove := true
			if warehouse[row+1][col] == wall || warehouse[row+1][col+1] == wall {
				return false
			}
			if warehouse[row+1][col] == boxL || warehouse[row+1][col] == boxR {
				canMove = canMove && canMoveBox(warehouse, row+1, col-(warehouse[row+1][col]-boxL), direction)
			}
			if warehouse[row+1][col+1] == boxL {
				canMove = canMove && canMoveBox(warehouse, row+1, col+1, direction)
			}
			return canMove
		}
	case "<":
		for j := col - 1; j >= 0 && warehouse[row][j] != wall; j-- {
			if warehouse[row][j] == 0 {
				return true
			}
		}
		return false
	case ">":
		for j := col + 2; j < len(warehouse[row]) && warehouse[row][j] != wall; j++ {
			if warehouse[row][j] == 0 {
				return true
			}
		}
		return false
	}
	return false
}

func moveBox(warehouse [][]int, row int, col int, direction string) {
	switch direction {
	case "^":
		if warehouse[row-1][col] == boxL || warehouse[row-1][col] == boxR {
			moveBox(warehouse, row-1, col-(warehouse[row-1][col]-boxL), direction)
		}
		if warehouse[row-1][col+1] == boxL {
			moveBox(warehouse, row-1, col+1, direction)
		}
		warehouse[row-1][col] = boxL
		warehouse[row-1][col+1] = boxR
		warehouse[row][col] = 0
		warehouse[row][col+1] = 0
	case "v":
		if warehouse[row+1][col] == boxL || warehouse[row+1][col] == boxR {
			moveBox(warehouse, row+1, col-(warehouse[row+1][col]-boxL), direction)
		}
		if warehouse[row+1][col+1] == boxL {
			moveBox(warehouse, row+1, col+1, direction)
		}
		warehouse[row+1][col] = boxL
		warehouse[row+1][col+1] = boxR
		warehouse[row][col] = 0
		warehouse[row][col+1] = 0
	case "<":
		j := col - 1
		for ; warehouse[row][j] != 0; j-- {
		}
		for ; j < col+1; j++ {
			warehouse[row][j] = warehouse[row][j+1]
			warehouse[row][j+1] = 0
		}
	case ">":
		j := col + 2
		for ; warehouse[row][j] != 0; j++ {
		}
		for ; j > col; j-- {
			warehouse[row][j] = warehouse[row][j-1]
			warehouse[row][j-1] = 0
		}
	}
}

func moveWithWideBoxes(move string, robot Robot, warehouse [][]int) Robot {
	switch move {
	case "^":
		if robot.row > 0 {
			if warehouse[robot.row-1][robot.col] == wall {
				return robot
			}
			if warehouse[robot.row-1][robot.col] == boxL || warehouse[robot.row-1][robot.col] == boxR {
				if canMoveBox(warehouse, robot.row-1, robot.col-(warehouse[robot.row-1][robot.col]-boxL), move) {
					moveBox(warehouse, robot.row-1, robot.col-(warehouse[robot.row-1][robot.col]-boxL), move)
				} else {
					return robot
				}
			}
			robot.row--
		}
	case "v":
		if robot.row < len(warehouse)-1 {
			if warehouse[robot.row+1][robot.col] == wall {
				return robot
			}
			if warehouse[robot.row+1][robot.col] == boxL || warehouse[robot.row+1][robot.col] == boxR {
				if canMoveBox(warehouse, robot.row+1, robot.col-(warehouse[robot.row+1][robot.col]-boxL), move) {
					moveBox(warehouse, robot.row+1, robot.col-(warehouse[robot.row+1][robot.col]-boxL), move)
				} else {
					return robot
				}
			}
			robot.row++
		}
	case "<":
		if robot.col > 0 {
			if warehouse[robot.row][robot.col-1] == wall {
				return robot
			}
			if warehouse[robot.row][robot.col-1] == boxR {
				if canMoveBox(warehouse, robot.row, robot.col-2, move) {
					moveBox(warehouse, robot.row, robot.col-2, move)
				} else {
					return robot
				}
			}
			robot.col--
		}
	case ">":
		if robot.col < len(warehouse[robot.row])-1 {
			if warehouse[robot.row][robot.col+1] == wall {
				return robot
			}
			if warehouse[robot.row][robot.col+1] == boxL {
				if canMoveBox(warehouse, robot.row, robot.col+1, move) {
					moveBox(warehouse, robot.row, robot.col+1, move)
				} else {
					return robot
				}
			}
			robot.col++
		}
	}
	return robot
}

func printWarehouse(warehouse [][]int, robot Robot) {
	for i := range warehouse {
		for j := range warehouse[i] {
			if i == robot.row && j == robot.col {
				fmt.Print("@")
			} else if warehouse[i][j] == wall {
				fmt.Print("#")
			} else if warehouse[i][j] == boxL {
				fmt.Print("[")
			} else if warehouse[i][j] == boxR {
				fmt.Print("]")
			} else if warehouse[i][j] == box {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
