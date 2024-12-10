package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

type Peak struct {
	row int
	col int
}

func scoreTrailhead(trailMap [][]int, row int, col int) []Peak {
	// fmt.Println(trailMap[row][col], ":", row, col)
	if trailMap[row][col] == 9 {
		// fmt.Println("Trail!")
		return []Peak{{row: row, col: col}}
	}

	trails := []Peak{}
	if row > 0 && trailMap[row-1][col] == trailMap[row][col]+1 {
		trails = append(trails, scoreTrailhead(trailMap, row-1, col)...)
	}
	if col < len(trailMap[row])-1 && trailMap[row][col+1] == trailMap[row][col]+1 {
		trails = append(trails, scoreTrailhead(trailMap, row, col+1)...)
	}
	if row < len(trailMap)-1 && trailMap[row+1][col] == trailMap[row][col]+1 {
		trails = append(trails, scoreTrailhead(trailMap, row+1, col)...)
	}
	if col > 0 && trailMap[row][col-1] == trailMap[row][col]+1 {
		trails = append(trails, scoreTrailhead(trailMap, row, col-1)...)
	}
	return trails
}

func dedupePeaks(peaks []Peak) []Peak {
	unique := []Peak{}
	hashes := map[int]bool{}

	for _, peak := range peaks {
		hash := peak.row*100 + peak.col
		if !hashes[hash] {
			unique = append(unique, peak)
			hashes[hash] = true
		}
	}
	return unique
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	trailMap := lib.Read2dArray(stdin, false)
	for i := range trailMap {
		for j := range trailMap[0] {
			if trailMap[i][j] == 0 {
				score := scoreTrailhead(trailMap, i, j)
				score = dedupePeaks(score)
				// fmt.Println("Trailhead at", i, j, "scores", len(score))
				result += len(score)

			}
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	trailMap := lib.Read2dArray(stdin, false)
	for i := range trailMap {
		for j := range trailMap[0] {
			if trailMap[i][j] == 0 {
				score := scoreTrailhead(trailMap, i, j)
				// fmt.Println("Trailhead at", i, j, "scores", len(score))
				result += len(score)

			}
		}
	}
	return fmt.Sprint(result)
}
