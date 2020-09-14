package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	var ret = [][]int{}

	var dfs func(r *TreeNode, level int)
	dfs = func(r *TreeNode, level int) {
		if r == nil {
			return
		}
		ret = append(ret, []int{})
		ret[level] = append(ret[level], r.Val)
		dfs(r.Left, level+1)
		dfs(r.Right, level+1)
	}
	dfs(root, 0)
	return ret
}
