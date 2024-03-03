package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var sums []int
	var sum, carryover, unit int
	curr1 := l1
	curr2 := l2
	for curr1 != nil || curr2 != nil {
		if curr1 == nil {
			sum = curr2.Val
		} else if curr2 == nil {
			sum = curr1.Val
		} else if curr1 != nil && curr2 != nil {
			sum = curr1.Val + curr2.Val
		}
		if carryover > 0 {
			sum += carryover
		}
		if sum > 9 {
			unit = sum % 10
			carryover = 1
			sum = unit
		}
		sums = append(sums, sum)
		if curr1.Next != nil {
			curr1 = curr1.Next
		}
		if curr2.Next != nil {
			curr2 = curr2.Next
		}
	}
	for _, n := range sums {
		addNewNode(res, n)
	}
	return res
}

func addNewNode(l *ListNode, v int) {
	newNode := new(ListNode)
	newNode.Val = v
	if l == nil {
		l = newNode
	} else {
		curr := l
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
}
