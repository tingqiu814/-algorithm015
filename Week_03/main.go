package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("maxSlidingWindow")
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println("replaceSpace")
	fmt.Println("expect We%20are%20happy. got: ", replaceSpace("We are happy."))
	//fmt.Println("combine")
	//fmt.Println(combine(4, 2))
	fmt.Println("reversePrint")
	input := Arr2List([]int{1, 3, 2})
	start := time.Now().UnixNano()
	fmt.Println("input: ", List2Arr(input), ", expect [2, 3, 1] , output: ", reversePrint(input))
	fmt.Println("during: ", GetCostTime(start))
}
func GetCostTime(start int64) int64 {
	end := time.Now().UnixNano()
	return (end - start) // / 1e6
}
func NewStack() *Stack {
	return &Stack{
		0,
		[]interface{}{},
	}
}

type Stack struct {
	len  int
	data []interface{}
}

func (c *Stack) Pop() interface{} {
	if c.len > 0 {
		var ret = c.data[c.len-1]
		c.data = c.data[:c.len-1]
		c.len = c.len - 1
		return ret
	}
	return nil
}
func (c *Stack) Push(item interface{}) {
	c.len = c.len + 1
	c.data = append(c.data, item)
}

func (c *Stack) Len() int {
	return c.len
}

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
