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

func isXmasRecurse(search [][]string, remaining string, row int, col int, rowDir int, colDir int) bool {
	if row < 0 || row >= len(search) || col < 0 || col >= len(search[row]) {
		return false
	}
	if search[row][col] != remaining[:1] {
		return false
	}
	if len(remaining) == 1 {
		return true
	}
	return isXmasRecurse(search, remaining[1:], row+rowDir, col+colDir, rowDir, colDir)
}

func xmasSpellCount(search [][]string, row int, col int) int {
	if search[row][col] != "X" {
		return 0
	}
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r != 0 || c != 0 {
				if isXmasRecurse(search, "MAS", row+r, col+c, r, c) {
					count++
				}
			}
		}
	}
	return count
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	wordSearch := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
	}

	for row := range wordSearch {
		for col := range wordSearch[row] {
			count := xmasSpellCount(wordSearch, row, col)
			if count > 0 {
				result += count
			}
		}
	}
	return fmt.Sprint(result)
}

func isXmasCross(search [][]string, row int, col int) bool {
	if search[row][col] != "A" {
		return false
	}
	corners := search[row-1][col-1] + search[row-1][col+1] + search[row+1][col-1] + search[row+1][col+1]
	return strings.Count(corners, "M") == 2 && strings.Count(corners, "S") == 2 && search[row-1][col-1] != search[row+1][col+1]
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	wordSearch := [][]string{}
	for stdin.Scan() {
		line := stdin.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
	}

	// out := lib.Create2dArray[string](len(wordSearch), len(wordSearch[0]), ".")

	for row := 1; row < len(wordSearch)-1; row++ {
		for col := 1; col < len(wordSearch[row])-1; col++ {
			if isXmasCross(wordSearch, row, col) {
				result++
				// fmt.Println(row, col)

				// out[row][col] = wordSearch[row][col]
				// out[row-1][col-1] = wordSearch[row-1][col-1]
				// out[row-1][col+1] = wordSearch[row-1][col+1]
				// out[row+1][col-1] = wordSearch[row+1][col-1]
				// out[row+1][col+1] = wordSearch[row+1][col+1]
			}
		}
	}
	// for _, l := range out {
	// 	fmt.Println(l)
	// }
	return fmt.Sprint(result)
}
