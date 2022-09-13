/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[

	1->4->5,
	1->3->4,
	2->6

]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []

Constraints:

k == lists.length
0 <= k <= 10 to the power of 4
0 <= lists[i].length <= 500
-10 to the power of 4 <= lists[i][j] <= 10 to the power of 4
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 10 to the power of 4.
*/
package main

import (
	//	"bufio"
	"fmt"
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

	l4 := new(ListNode)
	node := &ListNode{Val: 3}
	l4.Next = node
	node = &ListNode{Val: 4, Next: l4.Next}
	l4.Next = node
	node = &ListNode{Val: 2, Next: l4.Next}
	l4.Next = node

	l3 := new(ListNode)
	node = &ListNode{Val: 4}
	l3.Next = node
	node = &ListNode{Val: 6, Next: l3.Next}
	l3.Next = node
	node = &ListNode{Val: 5, Next: l3.Next}
	l3.Next = node

	l1 := new(ListNode)
	node = &ListNode{Val: 9}
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

	var listsN []*ListNode
	listsN = append(listsN, nil)

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

	l5 := mergeKLists(listsN)
	fmt.Println("Result")
	for l5 != nil {
		fmt.Println(l5.Val)
		l5 = l5.Next
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	mergeTwoLists := func(list1 *ListNode, list2 *ListNode) *ListNode {
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
	var listReturn *ListNode
	if len(lists) > 0 {
		listReturn = lists[0]
	} else {
		listReturn = nil
	}

	for i := 1; i <= len(lists)-1; i++ {
		listReturn = mergeTwoLists(listReturn, lists[i])
	}
	return listReturn
}
