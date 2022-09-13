package main

/*Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21


Constraints:

-2147483648<= x <= 2147483647
*/
import (
	"fmt"
	"strconv"
)

func main() {
	x := 1534236469
	z := reverse(x)
	fmt.Print(z)

}

func reverse(x int) int {
	plus := true
	if x < 0 {
		plus = false
		x = -x
	}
	strX := strconv.Itoa(x)
	var strR string
	for i := len(strX) - 1; i >= 0; i-- {
		strR += string(strX[i])
	}
	if !plus {
		strR = "-" + strR
	}

	rev, er := strconv.Atoi(strR)
	if er != nil {
		rev = 0
	}
	if rev < -2147483648 || rev > 2147483647 {
		rev = 0
	}
	return rev

}
