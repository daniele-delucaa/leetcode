package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
