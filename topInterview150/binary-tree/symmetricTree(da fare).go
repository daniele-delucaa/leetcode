package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fa la bfs per livelli e aggiunge un nodo 0 quando è nil
// infine controlla se i nodi "sono palindromi", se si,
// l'albero è simmetrico, altrimenti no
// la soluzione passa 196/199 testcase, non va bene
func isSymmetric(root *TreeNode) bool {
	var currLevel []TreeNode
	var allLevels [][]TreeNode
	currLevel = append(currLevel, *root)

	for len(currLevel) > 0 {
		allLevels = append(allLevels, currLevel)
		var nextLevel []TreeNode
		for _, n := range currLevel {
			if n.Left != nil {
				nextLevel = append(nextLevel, *n.Left)
			} else {
				nextLevel = append(nextLevel, TreeNode{})
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, *n.Right)
			} else {
				nextLevel = append(nextLevel, TreeNode{})
			}
		}
		currLevel = nextLevel
		if lastLevel(currLevel) {
			break
		}
	}
	if isPalindrome(allLevels) {
		return true
	}
	return false
}

func isPalindrome(allLevelsNodes [][]TreeNode) bool {
	for rowIndex, row := range allLevelsNodes {
		if rowIndex == 0 {
			continue
		}
		var j int = len(row) - 1
		for i := 0; i < len(row)/2; i++ {
			if row[i].Val != row[j].Val {
				return false
			}
			j--
		}
	}
	return true
}

// ritorna true se l'array contiene solo zeri, quindi siamo all'ultimo
// livello dell'albero
func lastLevel(level []TreeNode) bool {
	var zero bool = true
	for _, n := range level {
		if n.Val != 0 {
			zero = false
			break
		}
	}
	if !zero {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return dfs(root.Left, root.Right)
}

func dfs(l, r *TreeNode) bool {
	if l.Left == nil && r.Right == nil {
		return true
	}
	if l.Left == nil || r.Right == nil {
		return false
	}
	return (l.Val == r.Val && dfs(l.Left, r.Right) && dfs(l.Right, r.Left))
}
