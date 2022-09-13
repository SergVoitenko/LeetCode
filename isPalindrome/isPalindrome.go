package main

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.


Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1


Follow up: Could you solve it without converting the integer to a string?
*/

import (
	//	"bufio"
	"fmt"
	//	"os"
	//
	// "strings"
)

func main() {
	x := 9
	okPal := isPalindrome(x)
	fmt.Print(okPal)

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 9 {
		return true
	}
	var okPalimdrome bool
	cx := x
	var pal int
	for {
		modulo := cx % 10
		pal = pal*10 + modulo
		cx = cx / 10
		if cx < 10 {
			pal = pal*10 + cx
			break
		}
	}
	if pal == x {

		okPalimdrome = true
	}

	return okPalimdrome
}
