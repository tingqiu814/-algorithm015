package main

// 数学总结的方法，就是求n有几个5倍数
func trailingZeroes(n int) int {
	var ret = 0
	for n >= 5 {
		n = n / 5
		ret += n
	}
	return ret
}
func trailingZeroes2(n int) int {
	if n < 5 {
		return 0
	}
	var total = int64(1)
	for i := 1; i <= n; i++ {
		total = total * int64(i)
	}
	var ret = 0
	for total%10 == 0 {
		ret += 1
		total /= 10
	}
	return ret
}
