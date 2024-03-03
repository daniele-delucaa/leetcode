package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// soluzione in cui faccio visite inorder sui due alberi, se
// trovo un nodo nil appendo alle slice delle visite uno 0
// confronto le due slice ottenute dalle visite, se sono
// uguali, ritorno true, false altrimenti
// N.B. soluzione non ottimale a causa dell'append dello 0
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var vis1, vis2 []int
	inorder(p, &vis1)
	inorder(q, &vis2)
	if len(vis1) != len(vis2) {
		return false
	}
	for i := range vis1 {
		if vis1[i] != vis2[i] {
			return false
		}
	}
	return true
}

func inorder(r *TreeNode, visited *[]int) {
	if r == nil {
		return
	}
	if r.Left == nil {
		*visited = append(*visited, 0)
	}
	if r.Right == nil {
		*visited = append(*visited, 0)
	}
	inorder(r.Left, visited)
	*visited = append(*visited, r.Val)
	inorder(r.Right, visited)
}

// soluzione migliore, ricorsiva:
// se abbiamo due "alberi nulli", restituisce true, se solo
// uno dei due è nullo restituisce false, se i valori dei nodi
// che confronta sono diversi ritorna false, altrimenti fa le
// chiamate ricorsive di isSameTree sui sottoalberi sx e dx
// sui due alberi
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
