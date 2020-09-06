package main

import "fmt"

func main() {
	//fmt.Println("vim-go")
	//var ret = intersect([]int{1, 2, 2, 4}, []int{2, 2})
	//fmt.Println(ret)
	fmt.Println("expect 3 got: ", numDecodings("226"))
	fmt.Println(fizzBuzz(1))
	fmt.Println(addDigits(10))
	fmt.Println("inorder:")
	//var root = fmt.Println(inorderTraversal())
	//var in = []int{1, nil, 2, 3}
	//fmt.Println(arr2BariTree(in))

	println(isAnagram("anagram", "nagaram"))
	println("347 top-k-frequent-elements")
	input := struct {
		In     []int
		K      int
		Expect []int
	}{
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	}
	fmt.Println("input: ", input, "return :", topKFrequent(input.In, input.K))
	input = struct {
		In     []int
		K      int
		Expect []int
	}{
		[]int{1},
		1,
		[]int{1},
	}
	fmt.Println("input: ", input, "return :", topKFrequent(input.In, input.K))
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
