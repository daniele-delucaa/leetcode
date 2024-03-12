package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// prefix sum approach
// we store all the increasing sum (the prefix sum of a node)
// in a map, we associate all the sums with the corresponding
// ListNode. After we re-iterate through the
// list starting from the dummy node, and we find the next value
// through the map.
func removeZeroSumSublists(head *ListNode) *ListNode {
	// initialization
	m := make(map[int]*ListNode) // key = prefixSum, value = corresponding node
	dummy := &ListNode{0, head}
	m[0] = dummy
	sum := 0
	curr := head

	// store the sums
	for curr != nil {
		sum += curr.Val
		m[sum] = curr
		curr = curr.Next
	}

	// "re-linking" of the LinkedList
	sum = 0
	curr = dummy
	for curr != nil {
		sum += curr.Val
		curr.Next = m[sum].Next
		curr = curr.Next
	}
	return dummy.Next
}
