package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 方法1 stack
*/
func reversePrint1(head *ListNode) []int {
	var ret []int
	var stack = NewStack()
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}

	n := stack.Pop()
	for n != nil {
		ret = append(ret, n.(int))
		n = stack.Pop()
	}
	return ret
}

/**
方法2 暴力求解
*/
func reversePrint(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}
