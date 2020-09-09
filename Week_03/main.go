package main

import "fmt"

func main() {
	fmt.Println("maxSlidingWindow")
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println("replaceSpace")
	fmt.Println("expect We%20are%20happy. got: ", replaceSpace("We are happy."))
	//fmt.Println("combine")
	//fmt.Println(combine(4, 2))
}

type Stack struct {
	len  int
	data []interface{}
}

func (c Stack) Pop() interface{} {
	if c.len > 0 {
		c.len = c.len - 1
		return c.data[c.len-1]
	}
	return nil
}
func (c Stack) Push(item interface{}) {
	c.len = c.len + 1
	c.data = append(c.data, item)
}

func (c Stack) Len() int {
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
