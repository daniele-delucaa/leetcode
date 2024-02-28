package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

}

func dfs(r *TreeNode) []*TreeNode {
	visited := []*TreeNode{}
	if r == nil {
		return visited
	}
	return recurse(r, visited)
}

func recurse(r *TreeNode, visited []*TreeNode) []*TreeNode {
	visited = append(visited, r)

	if r.Left != nil {
		visited = recurse(r.Left, visited)
	}
	if r.Right != nil {
		visited = recurse(r.Right, visited)
	}
	return visited
}
