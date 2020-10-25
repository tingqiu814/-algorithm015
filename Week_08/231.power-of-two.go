package main

import "fmt"

func main() {
	var inputs = []int{0, 1, 2, 3, 4, 5, -1}
	for _, input := range inputs {
		fmt.Println("input ", input, "output ", isPowerOfTwo(input))
	}

	var (
		arr1 = []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
		arr2 = []int{2, 1, 4, 3, 9, 6}
	)
	fmt.Println(relativeSortArray(arr1, arr2))
}
func isPowerOfTwo(n int) bool {
	// 二进制计算 清空最低位1 n & (n-1)
	//  11000
	//& 10111
	//  10000
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
