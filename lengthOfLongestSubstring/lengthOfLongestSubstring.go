package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)

func main() {
	str := "asdfsadfsadfsdfsfsdfsdfdsafds 1234 - $ % ! @ #4567890xcxvvcbrteweuhjd"
	lenSubst := lengthOfLongestSubstring(str)
	fmt.Println(lenSubst)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	lenSubst := 1
	for i := 0; i <= len(s)-1; i++ {
		lenS := 1
		mapSub := map[string]bool{string(s[i]): true}
		for j := i + 1; j <= len(s)-1; j++ {
			if mapSub[string(s[j])] {
				break
			} else {
				mapSub[string(s[j])] = true
				lenS += 1
			}

		}
		if lenS > lenSubst {
			lenSubst = lenS
		}

	}
	return lenSubst

}
