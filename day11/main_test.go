package main

import (
	"bufio"
	"os"
	"testing"
)

const part1Expected = "55312"
const part1Answer = "183484"
const part2Expected = "65601038650482"
const part2Answer = "218817038947400"

func loadfile(filename string) *bufio.Scanner {
	file, _ := os.Open(filename)
	return bufio.NewScanner(file)
}

func Test_Part1_Sample(t *testing.T) {
	scanner := loadfile("sample.txt")
	actual := part1(scanner)
	if part1Expected != actual {
		t.Errorf(`Part 1 sample incorrect. Expected: %s, Actual: %s`, part1Expected, actual)
	}
}

func Test_Part2_Sample(t *testing.T) {
	scanner := loadfile("sample.txt")
	actual := part2(scanner)
	if part2Expected != actual {
		t.Errorf(`Part 2 sample incorrect. Expected: %s, Actual: %s`, part2Expected, actual)
	}
}

func Test_Part1_Answer(t *testing.T) {
	scanner := loadfile("input.txt")
	actual := part1(scanner)
	if part1Answer != actual {
		t.Errorf(`Part 1 answer incorrect. Expected: %s, Actual: %s`, part1Answer, actual)
	}
}

func Test_Part2_Answer(t *testing.T) {
	scanner := loadfile("input.txt")
	actual := part2(scanner)
	if part2Answer != actual {
		t.Errorf(`Part 2 answer incorrect. Expected: %s, Actual: %s`, part2Answer, actual)
	}
}
