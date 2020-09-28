package main

import "fmt"

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0}, 0))
	fmt.Println(fourSum([]int{1, 0, -1}, 0))

	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
