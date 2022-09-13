package main

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
•	nums1.length == m
•	nums2.length == n
•	0 <= m <= 1000
•	0 <= n <= 1000
•	1 <= m + n <= 2000
•	-106 <= nums1[i], nums2[i] <= 106
*/

import (
	//	"bufio"
	"fmt"
	//	"os"
	//
	// "strings"
)

func main() {
	num1 := []int{6, 8}
	num2 := []int{2, 4}
	//	num2 := []int{6, 8, 10, 11, 12, 13, 14, 15, 16}
	//num2 := []int{-1, 1, 2, 16, 17, 18}
	//num1 := []int{6, 8, 10, 11, 12, 13, 14, 15, 16}
	median := findMedianSortedArrays(num1, num2)
	fmt.Printf("Результат: %f\n", median)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	var even bool
	var valCenter, valCenter1 int
	if (len(nums1)+len(nums2))%2 == 0 {
		even = true
	}
	center := (len(nums1) + len(nums2)) / 2
	if !even {
		center += 1
	}
	var curPos, valPos, pos1, pos2 int

	var find, step bool
	for !find {
		if curPos < center {
			if pos1 <= len(nums1)-1 && pos2 <= len(nums2)-1 {
				if nums1[pos1] > nums2[pos2] {
					valPos = nums2[pos2]
					curPos += 1
					pos2 += 1
					continue
				} else {
					valPos = nums1[pos1]
					curPos += 1
					pos1 += 1
					continue
				}
			}
			if pos1 > len(nums1)-1 {
				pos2 += center - curPos - 1
				valPos = nums2[pos2]
				curPos = center
				pos2 += 1
				continue
			}
			if pos2 > len(nums2)-1 {
				pos1 += center - curPos - 1
				valPos = nums1[pos1]
				curPos = center
				pos1 += 1
				continue

			}
		} else {
			if even {
				if step {
					valCenter1 = valPos
					find = true
				} else {
					valCenter = valPos
					center += 1
					step = true
					continue
				}

			} else {
				valCenter = valPos
				find = true
			}

		}

	}
	if even {
		return (float64(valCenter) + float64(valCenter1)) / 2
	}
	return float64(valCenter)
}
