/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/
package main

import (
	//	"bufio"
	"fmt"
	"strings"
	//	"os"
	//
	// "strings"
)

func main() {
	strs := []string{"dog", "racecar", "car"}
	longPre := longestCommonPrefix(strs)
	fmt.Println(longPre)

}

func longestCommonPrefix(strs []string) string {
	var longPre string
	var numOfChar int
	ok := false
	for !ok {
		if len(strs[0])-1 < numOfChar {
			break
		}
		newlongPre := longPre + string(strs[0][numOfChar])
		okPrefix := true
		for _, s := range strs {
			if !strings.HasPrefix(s, newlongPre) {
				ok = true
				okPrefix = false
				break
			}
		}
		if okPrefix {
			longPre = newlongPre
			numOfChar++
		}

	}
	return longPre
}
