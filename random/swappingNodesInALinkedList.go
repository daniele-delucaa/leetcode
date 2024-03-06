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

// soluzione con i due indici, senza l'array di supp
// iteriamo col for iniziale fino all'arrivo del primo nodo
// in posizione k e salviamo un puntatore a quel nodo.
// inizializziamo un altro puntatore che alla fine dovrà
// puntare in posizione k a partire dalla fine della lista
// con il secondo for facciamo avanzare i due indici,
// (curr parte dalla posizione k, r dalla head) appena curr
// raggiunge la fine della lista, r sarà nella posizione
// k a partire dalla fine della lista, infine esegui lo swap
// grazie ai due puntatori
func swapNodes(head *ListNode, k int) *ListNode {
	curr := head
	for i := 0; i < k-1; i++ {
		curr = curr.Next
	}
	l := curr // puntatore al nodo in posizione k da swappare
	r := head
	for curr.Next != nil {
		r = r.Next
		curr = curr.Next
	}
	l.Val, r.Val = r.Val, l.Val
	return head
}
