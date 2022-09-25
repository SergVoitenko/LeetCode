/*Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.



Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
Example 2:

Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
Example 3:

Input: s = ""
Output: 0


Constraints:

0 <= s.length <= 3 * 10 to the extent 4
s[i] is '(', or ')'.
*/

package main

import (
	"fmt"
)

func main() {
	s := "((())))))))))))(((())))((((((()))))))()()()()"

	// nums := []int{0, 4, 3, 0} // 0
	// nums := []int{3, 2, 3} // 6
	// nums := []int{-10, -1, -18, -19} // -19
	out := longestValidParentheses(s)
	fmt.Print(out)
}

/*
	func longestValidParentheses(s string) int {
		if len(s) == 0 || len(s) == 1{
			return 0
		}
		var lComb int
		var left, right, maxL int




		return lComb
	}
*/

func longestValidParentheses(s string) int {
	st := []int{-1}
	var lComb int
	for i, bracket := range s {
		if bracket == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = append(st, i)
			} else {
				if i-st[len(st)-1] > lComb {
					lComb = i - st[len(st)-1]
				}
			}
		}
	}
	return lComb
}
