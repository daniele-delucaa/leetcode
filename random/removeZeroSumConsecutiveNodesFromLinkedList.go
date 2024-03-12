package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// prefix sum approach
func removeZeroSumSublists(head *ListNode) *ListNode {
	m := make(map[int]*ListNode) // key = prefixSum, value = corresponding node
	dummy := &ListNode{0, head}
	m[0] = dummy

	sum := 0
	curr := head

	for curr != nil {
		sum += curr.Val
		m[sum] = curr
		curr = curr.Next
	}

	sum = 0
	curr = dummy
	for curr != nil {
		sum += curr.Val
		curr.Next = m[sum].Next
		curr = curr.Next
	}
	return dummy.Next
}
