package main

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	var x = num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}
