package main

import (
	"bufio"
	"os"
	"testing"
)

const part1Expected = "480"
const part1Answer = "37901"
const part2Expected = "875318608908"
const part2Answer = "77407675412647"

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
