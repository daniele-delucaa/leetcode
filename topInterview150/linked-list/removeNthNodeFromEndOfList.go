package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// soluzione macchinosa, troppi if
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var listLength int
	curr := head
	for curr != nil {
		listLength++
		curr = curr.Next
	}
	if listLength == 1 {
		head = nil
		return head
	}
	if listLength == 2 && n == 1 {
		head.Next = nil
		return head
	}
	if listLength == 2 && n == 2 {
		head = head.Next
		return head
	}
	var counter int
	index := listLength - n
	if index == 0 {
		head = head.Next
		return head
	}
	curr = head
	for curr != nil {
		counter++
		if counter == index {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
	return head
}
