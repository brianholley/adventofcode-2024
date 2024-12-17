package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

func readProgram(stdin *bufio.Scanner) (int, int, int, []int) {
	registers := [3]int{}
	for i := 0; i < 3; i++ {
		stdin.Scan()
		registers[i], _ = strconv.Atoi(stdin.Text()[len("Register A: "):])
	}
	stdin.Scan()
	stdin.Scan()
	program := stdin.Text()[len("Program: "):]
	return registers[0], registers[1], registers[2], lib.ParseStringOfIntsDelimited(program, ",")
}

func comboOperand(operand int, A int, B int, C int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	} else if operand == 4 {
		return A
	} else if operand == 5 {
		return B
	} else if operand == 6 {
		return C
	}
	return 0
}

func part1(stdin *bufio.Scanner) string {
	output := []string{}
	A, B, C, program := readProgram(stdin)

	ip := 0
	for ip < len(program) {
		literalOp := program[ip+1]
		comboOp := comboOperand(program[ip+1], A, B, C)
		// fmt.Println("Ip", ip, "Opcode", program[ip], "Literal op", literalOp, "Combo op", comboOp)
		switch program[ip] {
		case 0: //adv
			A = A / (1 << comboOp)
		case 1: //bxl
			B = B ^ literalOp
		case 2: //bst
			B = comboOp % 8
		case 3: //jnz
			if A != 0 {
				ip = literalOp
				continue
			}
		case 4: //bxc
			B = B ^ C
		case 5: //out
			output = append(output, fmt.Sprint(comboOp%8))
		case 6: //bdv
			B = A / (1 << comboOp)
		case 7: //cdv
			C = A / (1 << comboOp)
		}
		ip += 2
	}
	return strings.Join(output, ",")
}

func part2(stdin *bufio.Scanner) string {
	result := 0
	return fmt.Sprint(result)
}
