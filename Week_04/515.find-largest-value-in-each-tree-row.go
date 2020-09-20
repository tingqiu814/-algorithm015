package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var ret = []int{}
	var dfs func(r *TreeNode, idx int)
	dfs = func(r *TreeNode, idx int) {
		if r == nil {
			return
		}
		if len(ret) == idx {
			ret = append(ret, r.Val)
		}
		if r.Val > ret[idx] {
			ret[idx] = r.Val
		}
		dfs(r.Left, idx+1)
		dfs(r.Right, idx+1)
	}
	dfs(root, 0)

	return ret
}
