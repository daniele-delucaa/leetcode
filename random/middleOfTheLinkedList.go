package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	curr := head
	l, r := curr, curr
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}
	return l
}
