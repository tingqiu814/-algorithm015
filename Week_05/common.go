package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arr2BariTree(in []int) *TreeNode {
	var ret = &TreeNode{}

	return ret

}

func Arr2List(arr []int) *ListNode {
	var ret *ListNode
	var forRet *ListNode

	for _, v := range arr {
		if forRet == nil {
			forRet = &ListNode{Val: v}
			ret = forRet
		} else {
			forRet.Next = &ListNode{Val: v}
			forRet = forRet.Next
		}
	}

	return ret
}
func List2Arr(node *ListNode) []int {
	var ret = make([]int, 0)
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}
