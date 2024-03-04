func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	next := head.Next
	for next != nil {
		if curr.Val == next.Val {
			curr.Next = next.Next
			next = next.Next
			continue
		}
		curr, next = next, next.Next
	}
	return head
}

// codice fatto da me
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