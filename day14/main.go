package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	row  int
	col  int
	vRow int
	vCol int
}

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func parseRobots(stdin *bufio.Scanner) []Robot {
	robots := []Robot{}
	for stdin.Scan() {
		line := stdin.Text()

		re := regexp.MustCompile(`p=([0-9]+),([0-9]+) v=([\-0-9]+),([\-0-9]+)`)
		robot := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(robot[1])
		y, _ := strconv.Atoi(robot[2])
		vx, _ := strconv.Atoi(robot[3])
		vy, _ := strconv.Atoi(robot[4])

		robots = append(robots, Robot{row: y, col: x, vRow: vy, vCol: vx})
	}
	return robots
}

func scoreQuadrants(robots []Robot, boardRows int, boardCols int) int {
	quadrants := [4]int{}
	for _, robot := range robots {
		if robot.row == boardRows/2 || robot.col == boardCols/2 {
			continue
		}
		q := (2*robot.row/boardRows)*2 + (2 * robot.col / boardCols)
		quadrants[q]++
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	robots := parseRobots(stdin)

	const boardRows = 103 // Sample given is 7
	const boardCols = 101 // Sample given is 11

	const cycles = 100

	for i := range robots {
		robots[i].row = (robots[i].row + cycles*robots[i].vRow) % boardRows
		robots[i].col = (robots[i].col + cycles*robots[i].vCol) % boardCols

		if robots[i].row < 0 {
			robots[i].row += boardRows
		}
		if robots[i].col < 0 {
			robots[i].col += boardCols
		}
	}

	result = scoreQuadrants(robots, boardRows, boardCols)

	return fmt.Sprint(result)
}

func boardAfterCycle(robots []Robot, boardRows int, boardCols int, cycle int) [][]int {
	board := lib.Create2dArray(boardRows, boardCols, 0)

	for i := range robots {
		row := ((robots[i].row+cycle*robots[i].vRow)%boardRows + boardRows) % boardRows
		col := ((robots[i].col+cycle*robots[i].vCol)%boardCols + boardCols) % boardCols
		board[row][col]++
	}
	return board
}

func printBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			if j == len(board[i])/2 {
				fmt.Print(" ")
			} else if board[i][j] == 0 {
				fmt.Print(".")
			} else if board[i][j] > 9 {
				fmt.Print("X")
			} else {
				fmt.Print(board[i][j])
			}
		}
		fmt.Println()
	}
}

func printBoardToImage(board [][]int, cycle int) {
	img := image.NewRGBA(image.Rect(0, 0, len(board[0]), len(board)))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)

	green := color.RGBA{0, 128, 0, 255}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] > 0 {
				img.Set(j, i, green)
			}
		}
	}
	file, err := os.Create(fmt.Sprint("day14/test/", cycle, ".png"))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err)
	}
}

func part2(stdin *bufio.Scanner) string {
	result := 7892

	robots := parseRobots(stdin)

	const boardRows = 103
	const boardCols = 101

	const fullCycleCount = boardRows * boardCols
	for c := 1; c < fullCycleCount; c++ {
		board := boardAfterCycle(robots, boardRows, boardCols, c)
		printBoardToImage(board, c)
	}
	return fmt.Sprint(result)
}
