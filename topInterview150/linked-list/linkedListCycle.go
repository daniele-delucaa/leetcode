package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	occ := make(map[*ListNode]int)
	curr := head
	for curr != nil {
		if occ[curr] > 1 {
			return true
		}
		occ[curr]++
		curr = curr.Next
	}
	return false
}
