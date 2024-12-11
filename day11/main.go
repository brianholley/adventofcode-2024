package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"math"
	"slices"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		if stones[i] == 0 {
			stones[i] = 1
		} else {
			digits := int(math.Log10(float64(stones[i]))) + 1
			if digits%2 == 0 {
				left := stones[i] / int(math.Pow10(digits/2))
				right := stones[i] % int(math.Pow10(digits/2))

				stones[i] = left
				stones = slices.Insert(stones, i+1, right)
				i++
			} else {
				stones[i] *= 2024
			}
		}
	}
	return stones
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		stones := lib.ParseStringOfIntsSpaceDelimited(line)

		for i := 1; i <= 25; i++ {
			stones = blink(stones)
		}
		result = len(stones)
	}
	return fmt.Sprint(result)
}

func blink2(stones map[int]int) map[int]int {
	endStones := map[int]int{}
	for val, count := range stones {
		if val == 0 {
			endStones[1] += count
		} else {
			digits := int(math.Log10(float64(val))) + 1
			if digits%2 == 0 {
				left := val / int(math.Pow10(digits/2))
				right := val % int(math.Pow10(digits/2))

				endStones[left] += count
				endStones[right] += count
			} else {
				v := val * 2024
				endStones[v] += count
			}
		}
	}
	return endStones
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		stonesList := lib.ParseStringOfIntsSpaceDelimited(line)

		stones := map[int]int{}
		for i := range stonesList {
			stones[stonesList[i]]++
		}
		// fmt.Println("Start:", stones)
		for i := 1; i <= 75; i++ {
			stones = blink2(stones)
			// fmt.Println(i, ":", stones)
		}

		for _, count := range stones {
			result += count
		}
	}
	return fmt.Sprint(result)
}
