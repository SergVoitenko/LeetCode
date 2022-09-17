/*
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
Example 2:

Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Constraints:

The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000

Follow-up: Can you solve the problem in O(1) extra memory space?
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

	l3 := reverseKGroup(l1.Next, 3)
	fmt.Println("Результат")
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head
	}
	var startPoint, startPointOld, endPoint, point, nextPoint, returnList *ListNode
	i := 1
	first := true
	Kok := false
	returnList = head
	for head != nil {
		if !Kok {
			nextPoint = head
			j := k
			for head != nil {
				if j == 0 {
					break
				}
				j--
				head = head.Next

			}
			if j == 0 {
				head = nextPoint
				Kok = true
			} else {
				break
			}

		}
		if i == 1 {
			startPoint = head

			endPoint = head
			nextPoint = head.Next
			head = nextPoint
		} else {
			point = head
			nextPoint = head.Next
			point.Next = endPoint
			startPoint.Next = nextPoint
			endPoint = point
			head = nextPoint
		}
		if i == k {
			if first {
				returnList = endPoint
				startPointOld = startPoint
				first = false
			} else {
				startPointOld.Next = endPoint
				startPointOld = startPoint
			}

			Kok = false
			i = 0
		}
		i++

	}

	return returnList
}
