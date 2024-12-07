package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func totalMatchesAddMultRecursive(total int, intermediate int, inputs []int) bool {
	if len(inputs) == 0 {
		return total == intermediate
	}
	return (totalMatchesAddMultRecursive(total, intermediate+inputs[0], inputs[1:]) ||
		totalMatchesAddMultRecursive(total, intermediate*inputs[0], inputs[1:]))
}

func canAddMultInputsForTotal(total int, inputs []int) bool {
	return totalMatchesAddMultRecursive(total, inputs[0], inputs[1:])
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		totalAndInputs := strings.Split(line, ": ")
		total, _ := strconv.Atoi(totalAndInputs[0])
		inputs := lib.ParseStringOfIntsSpaceDelimited(totalAndInputs[1])

		if canAddMultInputsForTotal(total, inputs) {
			// fmt.Println(line)
			result += total
		}

	}
	return fmt.Sprint(result)
}

func concat(a int, b int) int {
	return a*int(math.Pow10(int(math.Log10(float64(b)))+1)) + b
}

func totalMatchesAddMultConcatRecursive(total int, intermediate int, inputs []int) bool {
	if len(inputs) == 0 {
		return total == intermediate
	}
	return (totalMatchesAddMultConcatRecursive(total, intermediate+inputs[0], inputs[1:]) ||
		totalMatchesAddMultConcatRecursive(total, intermediate*inputs[0], inputs[1:]) ||
		totalMatchesAddMultConcatRecursive(total, concat(intermediate, inputs[0]), inputs[1:]))
}

func canAddMultIConcatnputsForTotal(total int, inputs []int) bool {
	return totalMatchesAddMultConcatRecursive(total, inputs[0], inputs[1:])
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()
		totalAndInputs := strings.Split(line, ": ")
		total, _ := strconv.Atoi(totalAndInputs[0])
		inputs := lib.ParseStringOfIntsSpaceDelimited(totalAndInputs[1])

		if canAddMultIConcatnputsForTotal(total, inputs) {
			// fmt.Println(line)
			result += total
		}
	}
	return fmt.Sprint(result)
}
