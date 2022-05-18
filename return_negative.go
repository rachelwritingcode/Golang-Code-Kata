// In this simple assignment you are given a number and have to make it negative.
// But maybe the number is already negative?
// Examples

// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0

// Notes
//     The number can be negative already, in which case no change is required.
//     Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

package kata

import "math"

func MakeNegative(x int) int {
	// Had to conver the int to float in order to convert it into a negative with the Signbit function
	num := float64(x)
	// Used Signbit from math package to determine if int was negative or not
	negative := math.Signbit(num)
	if negative == false {
		x *= -1
	}
	return x
}
