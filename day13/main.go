package main

import (
	"adventofcode2024/lib"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	result := lib.Run(part1, part2)
	fmt.Println(result)
}

type Button struct {
	x int
	y int
}

type Prize struct {
	x int
	y int
}

type Machine struct {
	a     Button
	b     Button
	prize Prize
}

func readMachines(stdin *bufio.Scanner) []Machine {
	machines := []Machine{}

	reButton := regexp.MustCompile(`X\+([0-9]+), Y\+([0-9]+)`)
	rePrize := regexp.MustCompile(`X\=([0-9]+), Y\=([0-9]+)`)
	for stdin.Scan() {
		buttonA := reButton.FindStringSubmatch(stdin.Text())
		aX, _ := strconv.Atoi(buttonA[1])
		aY, _ := strconv.Atoi(buttonA[2])
		stdin.Scan()

		buttonB := reButton.FindStringSubmatch(stdin.Text())
		bX, _ := strconv.Atoi(buttonB[1])
		bY, _ := strconv.Atoi(buttonB[2])
		stdin.Scan()

		prize := rePrize.FindStringSubmatch(stdin.Text())
		pX, _ := strconv.Atoi(prize[1])
		pY, _ := strconv.Atoi(prize[2])
		stdin.Scan()

		machines = append(machines, Machine{a: Button{x: aX, y: aY}, b: Button{x: bX, y: bY}, prize: Prize{x: pX, y: pY}})
	}
	return machines
}

func optimalButtonPressesBruteForce(machine Machine) (bool, int, int) {
	reached := false
	aReached := 0
	bReached := 0

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			x := a*machine.a.x + b*machine.b.x
			y := a*machine.a.y + b*machine.b.y
			if x == machine.prize.x && y == machine.prize.y {
				if !reached || a*3+b < aReached*3+bReached {
					aReached = a
					bReached = b
				}
				reached = true
			}
		}
	}

	return reached, aReached, bReached
}

func part1(stdin *bufio.Scanner) string {
	result := 0
	machines := readMachines(stdin)
	for _, machine := range machines {
		reached, a, b := optimalButtonPressesBruteForce(machine)
		if reached {
			fmt.Println(machine, a, b)
			result += 3*a + b
		}
	}
	return fmt.Sprint(result)
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func optimalButtonPresses(machine Machine) (bool, int, int) {
	// Solving the systems of equations:
	// 1.
	// A(ax) + B(bx) = px
	// A(ay) + B(by) = py

	// 2.
	// A = (px - B(bx)) / (ax)
	// A = (py - B(by)) / (ay)

	// 3. (ay)(px - B(bx)) = (ax)(py - B(by))
	// 4. (ay)(px) - B(ay)(bx) = (ax)(py) - B(ax)(by)
	// 5. B(ax)(by) - B(ay)(bx) = (ax)(py) - (ay)(px)
	// 6. B = ((ax)(py) - (ay)(px)) / ((ax)(by) - (ay)(bx))
	B := (machine.a.x*machine.prize.y - machine.a.y*machine.prize.x) / (machine.a.x*machine.b.y - machine.a.y*machine.b.x)
	A := (machine.prize.x - B*machine.b.x) / machine.a.x

	parity := Prize{x: machine.a.x*A + machine.b.x*B, y: machine.a.y*A + machine.b.y*B}

	if machine.prize.x == parity.x && machine.prize.y == parity.y {
		return true, A, B
	}

	return false, 0, 0
}

func part2(stdin *bufio.Scanner) string {
	const oops = 10000000000000

	result := 0
	machines := readMachines(stdin)
	for _, machine := range machines {
		machine.prize.x += oops
		machine.prize.y += oops

		reached, a, b := optimalButtonPresses(machine)
		if reached {
			// fmt.Println(machine, a, b)
			result += 3*a + b
		}
	}
	return fmt.Sprint(result)
}
