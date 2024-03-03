package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var v []int
	inorder(root, &v)
	//fmt.Println(&v)
	j := len(v) - 1
	// controllo se "visita inorder è palindroma", se non lo è,
	// l'albero non è simmetrico
	for i := 0; i < len(v)/2; i++ {
		if (v)[i] != (v)[j] {
			return false
		}
		j--
	}
	// controllo se i due sottoalberi sono uguali
	rootIndex := len(v) / 2
	subTree1 := v[:rootIndex]
	subTree2 := v[rootIndex+1:]
	if len(subTree1) == 1 && len(subTree2) == 1 {
		if subTree1[0] == subTree2[0] {
			return true
		}
	}
	for i := range subTree1 {
		if subTree1[i] != subTree2[i] {
			return true
		} else {
			return false
		}
	}
	return true
}

func inorder(r *TreeNode, visited *[]int) {
	if r == nil {
		return
	}
	inorder(r.Left, visited)
	*visited = append(*visited, r.Val)
	inorder(r.Right, visited)
}
