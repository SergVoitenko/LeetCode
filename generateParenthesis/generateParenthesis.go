/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8
*/
package main

import (
	"fmt"
)

func main() {
	p := generateParenthesis(9)
	fmt.Println(p)

}

func generateParenthesis(n int) []string {
	var addParent func(list *[]string, s string, l int, r int)
	addParent = func(list *[]string, s string, l int, r int) { // recursive function
		if l == 0 && r == 0 { //condition for terminating the recursion and adding the result to the slice
			*list = append(*list, s)
			return
		}

		if l > 0 { // add left parenthesis
			addParent(list, s+"(", l-1, r)
		}
		if l < r { // add right bracket
			addParent(list, s+")", l, r-1)
		}
	}
	var l []string
	addParent(&l, "", n, n)
	return l

}
