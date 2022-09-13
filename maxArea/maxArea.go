package main

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:
Input: height = [1,1]
Output: 1

Constraints:
•	n == height.length
•	2 <= n <= 105
•	0 <= height[i] <= 104
*/

import (
	//	"bufio"
	"fmt"
	//	"os"
	//
	// "strings"
)

func main() {
	height := []int{99, 1, 1, 100, 99, 1}
	maxAmount := maxArea(height)
	fmt.Println(maxAmount)

}
func maxArea(height []int) int {
	maxA := 0
	for poz, h := range height {
		var pLeft, pRight int
		for i := 0; i <= poz; i++ {
			if height[i] >= h {
				pLeft = i

				break
			}
		}
		for i := len(height) - 1; i >= poz; i-- {
			if height[i] >= h {
				pRight = i

				break
			}
		}
		volume := (pRight - pLeft) * h
		if volume > maxA {
			maxA = volume
		}

	}

	return maxA
}
