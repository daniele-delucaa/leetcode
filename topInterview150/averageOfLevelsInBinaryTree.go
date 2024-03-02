package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// uso l'array currLevel come se fosse una coda
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	currLevel := []*TreeNode{root}
	for len(currLevel) > 0 {
		sum := 0
		nextLevel := []*TreeNode{}
		for _, node := range currLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(currLevel)))
		currLevel = nextLevel
	}
	return res
}
