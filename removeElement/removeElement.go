/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.

	// It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums

	for (int i = 0; i < actualLength; i++) {
	    assert nums[i] == expectedNums[i];
	}

If all assertions pass, then your solution will be accepted.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{45, 67, 3, 2, 1, -50, 0, -100, -99, 0, 0, 0, 0, 1, 2, 3, 4, 44, 45, 61, 72, 8, 100, 11, 12, 12, 12, 12, 12}
	//	nums := []int{1, 1}
	p := removeElement(nums, 12)
	fmt.Println(p)

}
func removeElement(nums []int, val int) int {
	n := 0
	ok := false
	for !ok {
		for i := n; i <= len(nums)-1; i++ {
			if nums[i] == val {
				copy(nums[i:], nums[i+1:])
				nums = nums[:len(nums)-1]
				n = i
				break
			} else {
				n = len(nums)
			}

		}
		if n > len(nums)-1 {
			ok = true
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return len(nums)
}
