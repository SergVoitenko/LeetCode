/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:

Input: head = [1,2,3,4]
Output: [2,1,4,3]
Example 2:

Input: head = []
Output: []
Example 3:

Input: head = [1]
Output: [1]

Constraints:

The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/
package main

import (
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

	l1 := new(ListNode)
	node := &ListNode{Val: 9}
	l1.Next = node
	node = &ListNode{Val: 8, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 7, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 6, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 5, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 4, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 3, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 2, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 1, Next: l1.Next}
	l1.Next = node

	l3 := swapPairs(l1.Next)
	fmt.Println("Результат")
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var startPoint, startPointOld, endPoint, nextPoint, returnList *ListNode
	i := 1
	first := true
	returnList = head
	for head != nil {

		if i == 1 {
			startPoint = head
		} else {
			endPoint = head
			nextPoint = head.Next
			if first {
				returnList = endPoint
				endPoint.Next = startPoint
				startPoint.Next = nextPoint
				startPointOld = startPoint
				first = false
			} else {
				endPoint.Next = startPoint
				startPoint.Next = nextPoint
				startPointOld.Next = endPoint
				startPointOld = startPoint
			}

			head = startPoint
			i = 0

		}
		i++
		head = head.Next

	}

	return returnList
}
