package main

/* The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
*/
import (
	//	"bufio"
	"fmt"
	//	"os"
	//
	// "strings"
)

func main() {
	str := "PAYPALISHIRING"
	//str := "ABCD"
	num := 3
	str = convert(str, num)
	fmt.Println(str)

}

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	masStr := make([]string, numRows)
	endS := false
	indexS := 0
	startCol := true
	col := 0
	for !endS {
		for i := 0; i <= numRows-1; i++ {
			if indexS > len(s)-1 {
				endS = true
				break
			}
			if startCol {
				masStr[i] += string(s[indexS])
				indexS += 1
			} else {
				if i == numRows-1-col {
					masStr[i] += string(s[indexS])
					indexS += 1
				}
			}
		}
		if startCol {
			if numRows > 2 {
				col = 1
				startCol = false
			}

		} else {
			if col+1 == numRows-1 {
				col = 0
				startCol = true
			} else {
				col += 1
			}
		}

	}
	var convert string
	for _, st := range masStr {
		convert += st
	}
	return convert
}
