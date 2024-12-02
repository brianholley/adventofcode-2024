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

func part1(stdin *bufio.Scanner) string {
	result := 0

	left := []int{}
	right := []int{}

	for stdin.Scan() {
		line := stdin.Text()
		locationPair := lib.ParseStringOfIntsSpaceDelimited(line)
		left = append(left, locationPair[0])
		right = append(right, locationPair[1])
	}

	lib.ArraySortAscending(left)
	lib.ArraySortAscending(right)

	for i := range left {
		result = result + lib.Abs(left[i]-right[i])
	}

	return fmt.Sprint(result)
}

func generateFrequencyTable(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v] = freq[v] + 1
	}
	return freq
}

func part2(stdin *bufio.Scanner) string {
	result := 0

	left := []int{}
	right := []int{}

	for stdin.Scan() {
		line := stdin.Text()
		locationPair := lib.ParseStringOfIntsSpaceDelimited(line)
		left = append(left, locationPair[0])
		right = append(right, locationPair[1])
	}

	frequencies := generateFrequencyTable(right)

	for i := range left {
		result = result + (left[i] * frequencies[left[i]])
	}

	return fmt.Sprint(result)
}
