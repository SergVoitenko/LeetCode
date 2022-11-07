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

import "testing"

func Test_search(t *testing.T) {
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
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "Проверяем число которого нет в последовательности",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			want: -1,
		},
		{
			name: "Проверяем число большее чем в последовательности",
			args: args{nums: []int{1}, target: 0},
			want: -1,
		},
		{
			name: "Комбинация 1",
			args: args{nums: []int{6, 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: 0},
			want: 12,
		},
		{
			name: "Комбинация 2",
			args: args{nums: []int{6, 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: 10},
			want: 4,
		},
		{
			name: "Комбинация 3",
			args: args{nums: []int{6, 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: -10},
			want: 5,
		},
		{
			name: "Комбинация 4",
			args: args{nums: []int{6, 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: 8},
			want: 2,
		},
		{
			name: "Комбинация 5",
			args: args{nums: []int{6, 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: -7},
			want: -1,
		},
		{
			name: "Комбинация 6",
			args: args{nums: []int{ 7, 8, 9, 10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5}, target: 6},
			want: -1,
		},
		{
			name: "Комбинация 7",
			args: args{nums: []int{-9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5, 7, 8, 9, 10, -10 }, target: 6},
			want: -1,
		},
		{
			name: "Комбинация 7",
			args: args{nums: []int{10, -10, -9, -8, -6, -5, -4, -3, 0, 1, 2, 3, 4, 5, 7, 8, 9, }, target: 6},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
