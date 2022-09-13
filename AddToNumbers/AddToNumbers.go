package main

import (
	"fmt"
	"strconv"
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
	node = &ListNode{Val: 0, Next: l1.Next}
	l1.Next = node
	node = &ListNode{Val: 1, Next: l1.Next}
	l1.Next = node

	l2 := new(ListNode)
	node = &ListNode{Val: 8}
	l2.Next = node
	node = &ListNode{Val: 7, Next: l2.Next}
	l2.Next = node
	node = &ListNode{Val: 5, Next: l2.Next}
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

	l3 := addTwoNumbers(l1.Next, l2.Next)
	fmt.Println("Result")
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numList := func(l *ListNode) (numStr string) {
		for l != nil {
			numStr += strconv.Itoa(l.Val)

			l = l.Next
		}

		return
	}
	numStr1 := numList(l1)
	numStr2 := numList(l2)
	lStr := len(numStr1)
	if lStr < len(numStr2) {
		lStr = len(numStr2)
	}

	var sumStr string
	remnant := 0
	for i := 0; i <= lStr-1; i++ {
		num1 := 0
		num2 := 0
		if i <= len(numStr1)-1 {
			num1, _ = strconv.Atoi(string(numStr1[i]))
		}
		if i <= len(numStr2)-1 {
			num2, _ = strconv.Atoi(string(numStr2[i]))
		}
		sum := num1 + num2 + remnant
		remnant = sum / 10
		el := sum % 10
		sumStr += strconv.Itoa(el)
	}
	if remnant != 0 {
		sumStr += strconv.Itoa(remnant)
	}

	l3 := new(ListNode)
	for i := len(sumStr) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(sumStr[i]))
		node := &ListNode{Val: num}

		node.Next = l3.Next

		l3.Next = node
	}
	return l3.Next
}
