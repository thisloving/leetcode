package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	carry := 0

	head := &ListNode{Val: 0, Next: nil}
	list := head
	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}

		if l2 != nil {
			val2 = l2.Val
		}

		val := val1 + val2 + carry
		list.Val = val % 10
		carry = val / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		list.Next = &ListNode{Val: 0, Next: nil}
		list = list.Next
	}

	return head
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 9, Next: nil}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	list := addTwoNumbers(l1, l2)

	head := list
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
