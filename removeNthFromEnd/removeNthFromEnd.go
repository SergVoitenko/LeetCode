/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.



Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]


Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

package main

import (
	//	"bufio"
	"fmt"
	//"sort"
	//
	// "strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/*
		l1 := new(ListNode)
		node := &ListNode{Val: 1}
		l1.Next = node
		poz := 1
	*/

	l1 := new(ListNode)
	node := &ListNode{Val: 2}
	l1.Next = node
	node = &ListNode{Val: 1, Next: l1.Next}
	l1.Next = node

	poz := 2

	/*
		l1 := new(ListNode)
		node := &ListNode{Val: 5}
		l1.Next = node
		node = &ListNode{Val: 4, Next: l1.Next}
		l1.Next = node
		node = &ListNode{Val: 3, Next: l1.Next}
		l1.Next = node
		node = &ListNode{Val: 2, Next: l1.Next}
		l1.Next = node
		node = &ListNode{Val: 1, Next: l1.Next}
		l1.Next = node
		poz:=2
	*/

	l3 := removeNthFromEnd(l1.Next, poz)
	fmt.Println("Результат")
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	outHead := head
	mas := make(map[int]*ListNode)
	var amount int
	for head != nil {
		amount++
		mas[amount] = head
		head = head.Next
	}
	pointStart := amount - n
	if pointStart != 0 {
		pointNext := pointStart + 2
		head = mas[pointStart]
		if pointNext > amount {
			head.Next = nil
		} else {
			head.Next = mas[pointNext]
		}
	} else {
		if amount == 1 {
			outHead = nil
		} else {
			outHead = mas[2]
		}

	}

	return outHead
}
