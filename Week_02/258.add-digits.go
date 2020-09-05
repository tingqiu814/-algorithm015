package main

// 暴力求解
func addDigits(num int) int {
	for num >= 10 {
		num = num%10 + addDigits(num/10)
	}
	return num
}
