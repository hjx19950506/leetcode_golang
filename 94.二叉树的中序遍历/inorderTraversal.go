package inorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type S struct {
	root *TreeNode
	visited bool
}
//一种统一的方法实现迭代式的三种遍历
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	s := []S{}
	s = append(s, S{root, false})
	visited := false
	for len(s) != 0 {
		root = s[len(s)-1].root
		visited = s[len(s)-1].visited
		s = s[:len(s)-1]
		if root == nil {
			continue
		}
		if visited {
			res = append(res, root.Val)
		} else {
			s = append(s, S{root.Right, false})
			s = append(s, S{root, true})
			s = append(s, S{root.Left, false})
		}
	}
	return res
}