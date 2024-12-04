package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	for stdin.Scan() {
		line := stdin.Text()

		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		valid := re.FindAllString(line, -1)

		for _, v := range valid {
			parts := strings.Split(v[4:len(v)-1], ",")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			result += a * b
		}
	}
	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	enabled := true
	for stdin.Scan() {
		line := stdin.Text()

		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`)
		valid := re.FindAllString(line, -1)

		for _, v := range valid {
			if v == "do()" {
				// fmt.Println(v, enabled, "=>", true)
				enabled = true
			} else if v == "don't()" {
				// fmt.Println(v, enabled, "=>", false)
				enabled = false
			} else if enabled {
				parts := strings.Split(v[4:len(v)-1], ",")
				a, _ := strconv.Atoi(parts[0])
				b, _ := strconv.Atoi(parts[1])
				// fmt.Println(v, result, "=>", result+a*b)
				result += a * b
			} else {
				// fmt.Println(v, "skipped")
			}
		}
	}
	return fmt.Sprint(result)
}
