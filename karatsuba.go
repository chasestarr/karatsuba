package karatsuba

import (
	"math"
	"strconv"
	"strings"
)

// Multiply takes in two numbers and recursively multiplies them
func Multiply(a, b string) string {
	// base case, single digit number
	if len(a) < 2 || len(b) < 2 {
		intA, _ := strconv.Atoi(a)
		intB, _ := strconv.Atoi(b)
		return strconv.Itoa(intA * intB)
	}

	if len(a)%2 == 1 {
		a = "0" + a
	}
	if len(b)%2 == 1 {
		b = "0" + b
	}

	// added to fix uneven length factors
	a, b = addZeroFront(a, b)

	// split strings at middle index
	m := getMiddle(a, b)
	high1, low1 := splitNum(a, m)
	high2, low2 := splitNum(b, m)

	// recursive calls
	z0 := Multiply(low1, low2)
	z1 := Multiply(add(low1, high1), add(low2, high2))
	z2 := Multiply(high1, high2)

	t0 := sub(sub(z1, z2), z0)
	t1 := add(addZeroEnd(z0, m*2), addZeroEnd(t0, m))
	t2 := add(t1, z2)
	return trim(t2)
}

func getMiddle(a, b string) int {
	m := math.Max(float64(len(a)), float64(len(b)))
	return int(m / 2)
}

func splitNum(s string, m int) (string, string) {
	high := s[m:]
	low := s[:m]
	return high, low
}

func addZeroFront(a, b string) (string, string) {
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	return a, b
}

func addZeroEnd(s string, m int) string {
	return s + strings.Repeat("0", m)
}

func trim(s string) string {
	for s[:1] == "0" && len(s) > 1 {
		s = s[1:]
	}
	return s
}

func add(a, b string) string {
	var sum string
	var carry int
	a, b = addZeroFront(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		intA, _ := strconv.Atoi(a[i : i+1])
		intB, _ := strconv.Atoi(b[i : i+1])

		sumDigit := intA + intB + carry
		carry = 0

		if sumDigit >= 10 {
			sumDigit -= 10
			carry = 1
		}

		sum = strconv.Itoa(sumDigit) + sum
	}

	// handle carry after for-loop terminates
	if carry != 0 {
		return strconv.Itoa(carry) + sum
	}

	return sum
}

// does not work with negative number results
func sub(a, b string) string {
	var sum string
	var carry int
	a, b = addZeroFront(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		intA, _ := strconv.Atoi(a[i : i+1])
		intB, _ := strconv.Atoi(b[i : i+1])

		var sumDigit = intA - intB - carry
		carry = 0

		if sumDigit < 0 {
			sumDigit += 10
			carry = 1
		}

		sum = strconv.Itoa(sumDigit) + sum
	}
	return sum
}
