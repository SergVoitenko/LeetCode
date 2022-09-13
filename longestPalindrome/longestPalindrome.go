package main

import (
	//	"bufio"
	"fmt"
	//	"os"
	//
	// "strings"
)

func main() {

	//s := ""
	//s := "babad"
	//s := "babadadab"
	s := "cbbd"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	var center, radius int

	var s1 string
	for _, st := range s {
		s1 += "!" + string(st)
	}
	s1 += "!"

	bestC := 1
	bestR := 1

	for i := 2; i <= len(s1)-1; i++ {
		center = i
		radius = 0
		ok := false
		for !ok {
			if 0 <= center-radius-1 && center+radius+1 <= len(s1)-1 {
				if s1[center-radius-1] == s1[center+radius+1] {
					radius += 1
				} else {
					ok = true
				}
			} else {
				ok = true
			}
		}

		if radius > bestR {
			bestR = radius
			bestC = center
		}

	}

	s1 = string(s1[bestC-bestR : bestC+bestR+1])
	var palindrome string
	for _, st := range s1 {
		if string(st) != "!" {
			palindrome += string(st)
		}
	}
	return palindrome
}

/* func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	var center, radius int
	mapS := map[int]int{0: 0, 1: 1}
	var s1 string
	for _, st := range s {
		s1 += "!" + string(st)
	}
	s1 += "!"

	bestC := 1
	bestR := 1

	newC := 2
	for {
		for i := newC; i <= len(s1)-1; i++ {
			center = i
			radius = 0
			ok := false
			for !ok {
				if 0 <= center-radius-1 && center+radius+1 <= len(s1)-1 {
					if s1[center-radius-1] == s1[center+radius+1] {
						radius += 1
					} else {
						ok = true
					}
				} else {
					ok = true
				}
			}
			mapS[center] = radius
			if radius > bestR {
				bestR = radius
				bestC = center
			}
			if radius > 1 {
				delta := 0
				dm := 0
				for j := center - 1; j >= center-radius; j-- {
					dm += 1
					mapS[center+dm] = 0
					if mapS[j] > 1 {
						delta = center - j
						break
					}

				}
				if delta == 0 {
					newC = center + radius + 1
				} else {
					newC = center + delta
				}
				break
			}

		}
		if center >= len(s1)-1 || newC >= len(s1)-1 {
			break
		}

	}
	s1 = string(s1[bestC-bestR : bestC+bestR+1])
	var palindrome string
	for _, st := range s1 {
		if string(st) != "!" {
			palindrome += string(st)
		}
	}
	return palindrome
}*/
