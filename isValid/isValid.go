/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/
package main

import (
	//	"bufio"
	"fmt"
	//"strings"
	//	"sort"
	//
	// "strings"
)

func main() {
	s := "[[[{{}}]]](())[]{}[]())"
	tar := isValid(s)
	fmt.Println(tar)

}

func isValid(s string) bool {

	cut := func(s string, sep string) (before string, after string, found bool) {
		for i := 0; i <= len(s)-len(sep); i++ {
			found = true
			for j := 0; j <= len(sep)-1; j++ {
				if s[i+j] != sep[j] {
					found = false
					break
				}
			}
			if found {
				if i == 0 {
					after = s[len(sep):]
					break
				} else {
					if i == len(s)-len(sep) {
						before = s[0:i]
						break
					} else {
						before = s[0:i]
						after = s[i+len(sep):]
						break
					}
				}
			}
		}
		if !found {
			before = s
		}
		return before, after, found
	}

	for {
		left1, right1, ok1 := cut(s, "()")
		if ok1 {
			s = left1 + right1
			continue
		}
		left2, right2, ok2 := cut(s, "{}")
		if ok2 {
			s = left2 + right2
			continue
		}
		left3, right3, ok3 := cut(s, "[]")
		if ok3 {
			s = left3 + right3
			continue
		} else {
			break
		}

	}
	ok := true
	if len(s) > 0 {
		ok = false
	}

	return ok

}

/*
func isValid(s string) bool {
	for {
		left1, right1, ok1 := strings.Cut(s, "()")
		if ok1 {
			s = left1 + right1
			continue
		}
		left2, right2, ok2 := strings.Cut(s, "{}")
		if ok2 {
			s = left2 + right2
			continue
		}
		left3, right3, ok3 := strings.Cut(s, "[]")
		if ok3 {
			s = left3 + right3
			continue
		} else {
			break
		}

	}
	ok := true
	if len(s) > 0 {
		ok = false
	}

	return ok

}
*/
