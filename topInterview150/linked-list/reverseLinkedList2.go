package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var supp, supp2 []int
	curr := head
	var head2 *ListNode
	for curr != nil {
		supp = append(supp, curr.Val)
		curr = curr.Next
	}
	for i := range supp {
		if i >= left-1 && i < right {
			supp2 = append(supp2, supp[i])
		}
	}
	//supp2 = supp[left-1 : right]
	// ottenendo supp2 facendo subslicing,
	// l'istruzione supp[i] = supp2[j] fotte tutto,
	// si modifica supp, supp2 pure viene modificato
	// da quell'istruzione perche supp2 "punta" supp
	// perche supp2 è ottenuto facendo subslicing da supp
	// quindi è meglio iterare su supp e inserire
	// manualmente i valori su supp2
	j := len(supp2) - 1
	for i := range supp {
		if i >= left-1 && i < right {
			supp[i] = supp2[j]
			j--
		}
	}
	for i := 0; i < len(supp); i++ {
		if head2 == nil {
			head2 = new(ListNode)
			head2.Val = supp[i]
		} else {
			c := head2
			for c.Next != nil {
				c = c.Next
			}
			c.Next = new(ListNode)
			c.Next.Val = supp[i]
		}
	}
	return head2
}
