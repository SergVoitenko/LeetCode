/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]

Constraints:

1 <= nums.length <= 200
-10 в степени 9 <= nums[i] <= 10 в степени 9
-10 в степени 9 <= target <= 10 в степени 9
*/
package main

import (
	//	"bufio"
	"fmt"
	"sort"
	//
	// "strings"
)

func main() {
	//	nums := []int{1, 0, -1, 0, -2, 2}
	//	target := 0 // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	nums := []int{2, 2, 2, 2, 2}
	target := 8 // Output: [[2,2,2,2]]
	tar := fourSum(nums, target)
	fmt.Println(tar)

}

func fourSum(nums []int, target int) [][]int {
	mapList := make(map[int]int, len(nums))
	var findList [][]int

	for _, num := range nums {
		mapList[num]++

	}

	type key_value struct {
		key   int
		value int
	}

	var sorted_struct []key_value

	for key, value := range mapList {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].key < sorted_struct[j].key
	})

	/*	if target > 0 {
		if numMax < target {
			if numMax <= target/3 {
				posStart = len(sorted_struct) - 3
				if posStart < 0 {
					posStart = 0
				}
			} else {
				numStart = numMax / 3
				if numStart != 0 {
					firstStart = true
				}
			}
		}
	} */

	for i := 0; i <= len(sorted_struct)-1; i++ {
		k1 := sorted_struct[i]
		//		if firstStart && k1.key < numStart {
		//			continue
		//		}
		for j := i; j <= len(sorted_struct)-1; j++ {
			k2 := sorted_struct[j]
			if k2.key == k1.key && k2.value < 2 {
				continue
			}

			for k := j; k <= len(sorted_struct)-1; k++ {
				k3 := sorted_struct[k]
				if k3.key == k1.key && k3.value < 3 {
					continue
				}
				if k3.key == k2.key && k2.value < 2 {
					continue
				}
				for l := k; l <= len(sorted_struct)-1; l++ {
					k4 := sorted_struct[l]
					if k4.key == k1.key && k4.value < 4 {
						continue
					}
					if k4.key == k2.key && k4.value < 3 {
						continue
					}
					if k4.key == k3.key && k4.value < 2 {
						continue
					}
					sum := k1.key + k2.key + k3.key + k4.key
					if sum == target {
						el := []int{k1.key, k2.key, k3.key, k4.key}
						findList = append(findList, el)
					}
				}

			}

		}
	}

	return findList
}
