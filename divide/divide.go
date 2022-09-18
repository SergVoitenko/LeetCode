/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−2to the extent31, 2to the extent31 − 1]. For this problem, if the quotient is strictly greater than 2to the extent31 - 1, then return 2to the extent31 - 1, and if the quotient is strictly less than -2to the extent31, then return -2to the extent31.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.

Constraints:

-2 to the extent 31 <= dividend, divisor <= 2 to the extent 31 - 1
divisor != 0
*/
package main

import (
	"fmt"
)

func main() {

	p := divide(-2147483649, -1)
	fmt.Println(p)

}
func divide(dividend int, divisor int) int {
	var r int
	var dividendPlus, divisorPlus, oK bool
	if dividend >= 0 {
		dividendPlus = true
	} else {
		dividend = 0 - dividend
	}

	if divisor >= 0 {
		divisorPlus = true
	} else {
		divisor = 0 - divisor
	}
	if divisor == 1 {
		r = dividend
		oK = true

	}
	for !oK {
		if dividend >= divisor {
			dividend -= divisor
			r++
		} else {
			oK = true
		}

	}
	if (dividendPlus && !divisorPlus) || (!dividendPlus && divisorPlus) {
		r = 0 - r
	}
	if r < -2147483648 {
		r = -2147483648
	}
	if r > 2147483647 {
		r = 2147483647
	}
	return r
}
