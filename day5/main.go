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

func parseRules(stdin *bufio.Scanner) map[int][]int {
	rules := make(map[int][]int)
	for stdin.Scan() {
		line := stdin.Text()
		if line == "" {
			break
		}

		rule := lib.ParseStringOfIntsDelimited(line, "|")
		rules[rule[1]] = append(rules[rule[1]], rule[0])
	}
	return rules
}

func isPageCorrect(page int, update []int, rules map[int][]int, printed []int) bool {
	for _, r := range rules[page] {
		if !lib.ArrayContains(printed, r) && lib.ArrayContains(update, r) {
			return false
		}
	}
	return true
}

func isUpdateCorrect(update []int, rules map[int][]int) bool {
	printed := []int{}
	for _, u := range update {
		if !isPageCorrect(u, update, rules, printed) {
			return false
		}
		printed = append(printed, u)
	}
	return true
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	rules := parseRules(stdin)
	for stdin.Scan() {
		line := stdin.Text()

		update := lib.ParseStringOfIntsDelimited(line, ",")
		correct := isUpdateCorrect(update, rules)

		if correct {
			// fmt.Println("Correct! Adding ", update[len(update)/2])
			result += update[len(update)/2]
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	rules := parseRules(stdin)
	for stdin.Scan() {
		line := stdin.Text()

		update := lib.ParseStringOfIntsDelimited(line, ",")
		correct := isUpdateCorrect(update, rules)

		if !correct {
			toSort := lib.ArrayCopy(update)
			corrected := []int{}

			// fmt.Print("Correcting", toSort, "=> ")
			for len(toSort) > 0 {
				for i := range toSort {
					if isPageCorrect(toSort[i], update, rules, corrected) {
						corrected = append(corrected, toSort[i])
						toSort = lib.ArrayRemoveIndex(toSort, i)
						break
					}
				}
			}
			// fmt.Println(corrected)

			result += corrected[len(corrected)/2]
		}
	}
	return fmt.Sprint(result)
}
