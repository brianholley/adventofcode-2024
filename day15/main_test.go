package main

import (
	"bufio"
	"os"
	"testing"
)

var part1Samples = []string{"10092", "2028"}

const part1Answer = "1476771"

var part2Samples = []string{"9021"}

const part2Answer = "1468005"

func loadfile(filename string) *bufio.Scanner {
	file, _ := os.Open(filename)
	return bufio.NewScanner(file)
}

func Helper(t *testing.T, fn func(*bufio.Scanner) string, filename string, expected string) {
	scanner := loadfile(filename)
	actual := fn(scanner)
	if expected != actual {
		t.Errorf(`Incorrect. Expected: %s, Actual: %s`, expected, actual)
	}
}

func Test_Part1_Sample(t *testing.T) {
	Helper(t, part1, "sample.txt", part1Samples[0])
}

func Test_Part1_Sample2(t *testing.T) {
	Helper(t, part1, "sample2.txt", part1Samples[1])
}

func Test_Part2_Sample(t *testing.T) {
	Helper(t, part2, "sample.txt", part2Samples[0])
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
