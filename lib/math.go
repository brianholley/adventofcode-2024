package lib

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a int, b int) int {
	return (a * b) / Gcd(a, b)
}
