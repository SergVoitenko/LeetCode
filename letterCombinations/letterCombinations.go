/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]

Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
package main

import (
	//	"bufio"
	"fmt"
	//	"sort"
	//
	// "strings"
)

func main() {
	digitsTelNumber := "9"
	digitsTelNumSt := letterCombinations(digitsTelNumber)
	fmt.Println(digitsTelNumSt)

}

func letterCombinations(digits string) []string {
	var digComb []string
	if len(digits) == 0 {
		return digComb
	}
	numDisplay := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	num := numDisplay[string(digits[0])]
	for _, s := range num {
		digComb = append(digComb, string(s))
	}
	for i := 1; i <= len(digits)-1; i++ {
		ldigComb := len(digComb)
		for j := range digComb {
			num = numDisplay[string(digits[i])]
			for _, s := range num {
				digComb = append(digComb, digComb[j]+string(s))
			}
		}
		digComb = digComb[ldigComb:]

	}

	return digComb

}
