package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)

func main() {
	target := 6

	// nums := []int{0, 4, 3, 0} // 0
	nums := []int{3, 2, 3} // 6
	// nums := []int{-10, -1, -18, -19} // -19
	out := twoSum(nums, target)
	fmt.Print(out)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		dif := target - num
		bal, ok := m[dif]
		if ok {
			return []int{bal, i}
		}
		m[num] = i
	}
	return []int{}
}
