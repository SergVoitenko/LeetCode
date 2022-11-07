/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1


Constraints:

1 <= nums.length <= 5000
-10 to the extent 4 <= nums[i] <= 10 to the extent 4
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-10 to the extent 4 <= target <= 10 to the extent 4
Accepted
1,792,930





*/

package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 3}
	target := 0
	position := search(nums, target)
	fmt.Print(position)
}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if left == right {
		return -1
	}

	if nums[left] > nums[right] {
		var endPos int

		find := false
		if nums[left] > nums[left+1] {
			endPos = left
			find = true
		}
		
		for !find {
			endPos = (left + right) / 2

			if nums[endPos] > nums[endPos-1] && nums[endPos] > nums[endPos+1] {
				find = true
			} else {
				if nums[endPos] < nums[right] {
					right = endPos
				} else {
					left = endPos
				}
			}
		}

		if target < nums[endPos+1] {
			return -1
		}
		if target > nums[endPos] {
			return -1
		}
		if target > nums[0] {
			left = 0
			right = endPos
		} else {
			left = endPos + 1
			right = len(nums) - 1
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if left == right {
		return -1
	}
	if nums[left] > target {
		return -1
	}
	if nums[right] < target {
		return -1
	}
	var targetPos int
	find := false
	for !find {
		targetPos = (left + right) / 2
		if nums[targetPos] == target {
			find = true
		} else {
			if left+1 == right {
				targetPos = -1
				find = true
			} else {

				if nums[targetPos] < target {
					left = targetPos
				} else {
					right = targetPos
				}

			}
		}

	}
	return targetPos
}
