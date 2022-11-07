/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 10 to the extent 4
-10 to the extent 4 <= nums[i] <= 10 to the extent 4
nums contains distinct values sorted in ascending order.
-10 to the extent 4 <= target <= 10 to the extent 4



*/

package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 3, 5, 6}
	target := 5
	position := searchInsert(nums, target)
	fmt.Print(position)
}

func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	if nums[left] == target {
		return left
	}
	if nums[left] > target {
		return 0
	}
	if nums[right] == target {
		return right
	}
	if nums[right] < target {
		return right + 1
	}
	find := false
	var targetPos int
	for !find {
		targetPos = (left + right) / 2
		if nums[targetPos] == target {
			find = true
		} else {
			if left+1 == right {
				targetPos = left + 1
				find = true
			}
			if nums[targetPos] > target {
				right = targetPos
			} else {
				left = targetPos
			}
		}

	}
	return targetPos
}
