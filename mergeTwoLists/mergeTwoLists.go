/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
package main

import (
	//	"bufio"
	"fmt"
	//	"sort"
	//
	// "strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Input: l1 = [2,4,3], l2 = [5,6,4]
	// Output: [7,0,8]
	// Explanation: 342 + 465 = 807.
	/*l1 := new(ListNode)
	node := &ListNode{Val: 3}
	l1.Next = node
	node = &ListNode{Val: 4, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 2, Next: l1.Next}
	l1.Next = node

	l2 := new(ListNode)
	node = &ListNode{Val: 4}
	l2.Next = node
	node = &ListNode{Val: 6, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 5, Next: l2.Next}
	l2.Next = node*/

	l1 := new(ListNode)
	node := &ListNode{Val: 9}
	l1.Next = node
	node = &ListNode{Val: 3, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 2, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 1, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 0, Next: l1.Next}
	l1.Next = node

	l2 := new(ListNode)
	node = &ListNode{Val: 11}
	l2.Next = node
	node = &ListNode{Val: 10, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 8, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 2, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 1, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 0, Next: l2.Next}
	l2.Next = node

	/*l1 := new(ListNode)
	node := &ListNode{Val: 0}
	l1.Next = node

	l2 := new(ListNode)
	node = &ListNode{Val: 0}
	l2.Next = node */

	/*l1 := new(ListNode)
	node := &ListNode{Val: 9}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 9, Next: l1.Next}
	l1.Next = node

	l2 := new(ListNode)
	node = &ListNode{Val: 9}
	l2.Next = node
	node = &ListNode{Val: 9, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 9, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 9, Next: l2.Next}
	l2.Next = node */

	l3 := mergeTwoLists(l1.Next, l2.Next)
	fmt.Println("Результат")
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var list3 *ListNode
	nextL1 := false
	nextL2 := false
	if list1.Val < list2.Val {
		list3 = list1
		nextL1 = true

	} else {
		list3 = list2
		nextL2 = true
	}
	for nextL1 || nextL2 {
		if nextL1 {
			if list1.Next != nil {

				if list1.Next.Val < list2.Val {
					list1 = list1.Next
				} else {
					oldlist1 := list1.Next
					list1.Next = list2
					list1 = oldlist1
					nextL1 = false
					nextL2 = true
				}
			} else {
				list1.Next = list2
				nextL2 = false
				nextL1 = false
			}

		}
		if nextL2 {
			if list2.Next != nil {
				if list2.Next.Val < list1.Val {
					list2 = list2.Next
				} else {
					oldlist2 := list2.Next
					list2.Next = list1
					list2 = oldlist2
					nextL1 = true
					nextL2 = false
				}
			} else {
				list2.Next = list1
				nextL2 = false
				nextL1 = false
			}

		}
	}

	return list3
}
