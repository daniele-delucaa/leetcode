package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var supp []int
	curr := head
	for curr != nil {
		supp = append(supp, curr.Val)
		curr = curr.Next
	}
	supp[k-1], supp[len(supp)-k] = supp[len(supp)-k], supp[k-1]
	curr = head
	i := 0
	for curr != nil {
		if curr.Val != supp[i] {
			curr.Val = supp[i]
		}
		i++
		curr = curr.Next
	}
	return head
}
