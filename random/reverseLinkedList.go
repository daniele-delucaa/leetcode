package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var supp []int
	var head2 *ListNode
	curr := head
	for curr != nil {
		supp = append(supp, curr.Val)
		curr = curr.Next
	}
	for i := len(supp) - 1; i >= 0; i-- {
		if head2 == nil {
			head2 = new(ListNode)
			head2.Val = supp[i]
		} else {
			curr := head2
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = new(ListNode)
			curr.Next.Val = supp[i]
		}
	}
	return head2
}
