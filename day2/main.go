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

func areLevelsSafe(levels []int) bool {
	if levels[1] > levels[0] { // increasing
		for i := 1; i < len(levels); i++ {
			delta := levels[i] - levels[i-1]
			if delta < 1 || delta > 3 {
				return false
			}
		}
	} else { // decreasing
		for i := 1; i < len(levels); i++ {
			delta := levels[i] - levels[i-1]
			if delta < -3 || delta > -1 {
				return false
			}
		}
	}
	return true
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		report := stdin.Text()
		levels := lib.ParseStringOfIntsSpaceDelimited(report)

		if areLevelsSafe(levels) {
			result++
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		report := stdin.Text()
		levels := lib.ParseStringOfIntsSpaceDelimited(report)

		if areLevelsSafe(levels) {
			result++
		} else {
			for i := range levels {
				copy := lib.ArrayCopy(levels)
				copy = lib.ArrayRemoveIndex(copy, i)
				if areLevelsSafe(copy) {
					result++
					break
				}
			}
		}
	}
	return fmt.Sprint(result)
}
