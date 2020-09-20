package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("maxArea")
	fmt.Println("need 49 , got ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	obj := Constructor(10)
	for i := 0; i < 3; i++ {
		println(i, ":", obj.InsertFront(i))
	}
	param_1 := obj.InsertFront(100)
	param_2 := obj.InsertLast(99)
	param_3 := obj.DeleteFront()
	param_4 := obj.DeleteLast()
	param_5 := obj.GetFront()
	param_6 := obj.GetRear()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsFull()
	println(param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8)

	println("lemonadeChange")
	println("need true: ", lemonadeChange([]int{5, 5, 5, 10, 20}))
	println("need true: ", lemonadeChange([]int{5, 5, 10}))
	println("need false: ", lemonadeChange([]int{10, 10}))
	println("need false: ", lemonadeChange([]int{5, 5, 10, 10, 20}))

	println("numIslands")
	println("expect 1: ", numIslands([][]byte{[]byte("11110"), []byte("11010"), []byte("11000"), []byte("00000")}))
	println("expect 3: ", numIslands([][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011")}))
	println("expect 1: ", numIslands([][]byte{
		[]byte("111"),
		[]byte("010"),
		[]byte("111")}))

	println("trailingZeroes")
	println(trailingZeroes(30))

	println("search")
	println("expect 4 got: ", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

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
