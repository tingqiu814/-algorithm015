package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
//最近公共祖先，从root递归循环下来，只要找到root=任意目标，说明自己就是公共祖先， 返回root
//如果全小于root， 那说明两个都在左子树，递归左子树
//如果全大于root， 说明在右子树， 递归右子树
//如果一个再左一个在右，那肯定是公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if (root.Val < p.Val && root.Val > q.Val) || (root.Val > p.Val && root.Val < q.Val) {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}
