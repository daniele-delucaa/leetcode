package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// soluzione decente, sono autistico
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var listLength int
	curr := head
	for curr != nil {
		listLength++
		curr = curr.Next
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

// soluzione con due indici l e r, serve un nodo "dummy", che
// inizializziamo in testa alla lista (prima di head), facciamo partire
// i due indici l e r rispettivamente da dummy e head.
// il primo ciclo serve a far arrivare l'indice r in un nodo.
// dopo facciamo scorrere contemporaneamente l e r e appena
// r raggiunge la fine della lista (nil), l avrà raggiunto il
// nodo prima del nodo da cancellare, e cancelliamo il nodo dopo.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	l, r := dummy, head

	for i := 0; i < n; i++ {
		r = r.Next
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}
	// cancellazione nodo
	l.Next = l.Next.Next
	return dummy.Next
}
