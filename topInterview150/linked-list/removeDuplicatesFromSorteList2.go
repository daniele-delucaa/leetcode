package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var l, r *ListNode = curr, curr.Next
	for r != nil {
		if l.Val == r.Val {
			curr.Next = curr.Next.Next
			r = r.Next
			continue
		}
		curr = curr.Next
		l = l.Next
		r = r.Next
	}
	return head
}
