// Package goarmstrong provides function for checking Armstrong numbers.
package goarmstrong

import "math"

//IsArmstrong returns true if number is an armstrong number, false otherwise.
func IsArmstrong(number int64) bool {

	var remainder, result, n int64 = 0, 0, 0

	originalNumber := number
	for originalNumber != 0 {
		originalNumber /= 10
		n++
	}

	originalNumber = number
	for originalNumber != 0 {
		remainder = originalNumber % 10
		result += int64(math.Pow(float64(remainder), float64(n)))
		originalNumber /= 10
	}

	return result == number
}
