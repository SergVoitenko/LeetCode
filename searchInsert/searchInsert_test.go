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

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Проверяем число в последовательности",
			args: args{nums: []int{1, 3, 5, 6}, target: 5},
			want: 2,
		},
		{
			name: "Проверяем число которого нет в последовательности",
			args: args{nums: []int{1, 3, 5, 6}, target: 2},
			want: 1,
		},
		{
			name: "Проверяем число большее чем в последовательности",
			args: args{nums: []int{1, 3, 5, 6}, target: 7},
			want: 4,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
