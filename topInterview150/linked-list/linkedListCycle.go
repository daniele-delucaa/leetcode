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

// soluzione con algoritmo di lepre e tartaruga di floyd
// usare due indici (slow = tartaruga, fast = lepre), che partono
// dallo head della linkedList, iteriamo sulla linkedList
// e facciamo avanzare l'indice s di una posizione,
// l'indice f di due posizioni. L' indice f raggiungerà la
// fine della lista prima rispetto a s, se con s raggiungiamo una fine
// (senza ciclo) possiamo ritornare false, altrimenti se l'ultimo
// nodo sulla lista punta a un'altro elemento dentro la lista,
// si creerà un ciclo e restituiremo true quando i due indici s e f
// si incontreranno nello stesso nodo (se c'è un ciclo, si incontreranno
// sicuramente)
func hasCycle(head *ListNode) bool {
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if f == s {
			return true
		}
	}
	return false
}
