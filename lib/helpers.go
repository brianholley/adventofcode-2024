package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type PartFn func(*bufio.Scanner) string

func Run(part1 PartFn, part2 PartFn) string {
	if len(os.Args[1:]) == 0 {
		panic("Expected part number argument")
	}

	scanner := bufio.NewScanner(os.Stdin)

	switch part := os.Args[1]; part {
	case "1":
		return part1(scanner)
	case "2":
		return part2(scanner)
	default:
		panic("Unknown part: " + part)
	}
}

type Pos struct {
	Row int
	Col int
}

func Read2dArray(scanner *bufio.Scanner, spaceDelimited bool) [][]int {
	array := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if spaceDelimited {
			s := strings.Split(line, " ")
			row := make([]int, len(s))
			for i, v := range s {
				row[i], _ = strconv.Atoi(v)
			}
			array = append(array, row)
		} else {
			row := make([]int, len(line))
			for i, v := range line {
				row[i], _ = strconv.Atoi(string(v))
			}
			array = append(array, row)
		}
	}
	return array
}

func Create2dArray(rows int, cols int, defaultValue int) [][]int {
	arr := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			arr[i][j] = defaultValue
		}
	}
	return arr
}

func ParseStringOfIntsSpaceDelimited(str string) []int {
	numbers := []int{}
	s := strings.Split(str, " ")
	for _, v := range s {
		if len(v) > 0 {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}
	}
	return numbers
}
