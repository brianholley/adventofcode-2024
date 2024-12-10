package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
)

type Node struct {
	row int
	col int
}

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func findAntinodes(a Node, b Node) []Node {
	dRow := a.row - b.row
	dCol := a.col - b.col

	aa := Node{row: a.row + dRow, col: a.col + dCol}
	ab := Node{row: b.row - dRow, col: b.col - dCol}

	// fmt.Println(a, b, dRow, dCol, aa, ab)

	return []Node{aa, ab}
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	antennas := map[string][]Node{}
	i := 0
	mapWidth := 0
	for ; stdin.Scan(); i++ {
		line := stdin.Text()
		for j := range line {
			if line[j] != '.' {
				freq := string(line[j])
				antennas[freq] = append(antennas[freq], Node{row: i, col: j})
			}
		}
		mapWidth = len(line)
	}
	mapHeight := i

	antinodes := []Node{}

	for freq := range antennas {
		for a := range antennas[freq] {
			for b := a + 1; b < len(antennas[freq]); b++ {
				antinodes = append(antinodes, findAntinodes(antennas[freq][a], antennas[freq][b])...)
			}
		}
	}

	// De-dupe
	hashes := map[int]bool{}
	for _, node := range antinodes {
		if node.row < 0 || node.col < 0 || node.row >= mapHeight || node.col >= mapWidth {
			// fmt.Println("Antinode out of bounds:", node)
			continue
		}
		hash := node.row*100 + node.col
		if !hashes[hash] {
			// fmt.Println("Unique antinode found:", node)
			hashes[hash] = true
			result++
		} else {
			// fmt.Println("Duplicate antinode found:", node)
		}
	}

	return fmt.Sprint(result)
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
