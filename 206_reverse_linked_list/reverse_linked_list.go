package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Input: head = [1,2,3,4,5]
 * Output: [5,4,3,2,1]
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pre := &ListNode{}
	pre = nil
	next := &ListNode{}
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	revList := reverseList(list)
	for revList != nil {
		fmt.Printf("%d ", revList.Val)
		revList = revList.Next
	}
}
