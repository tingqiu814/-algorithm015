package main

func waysToStep(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}
	var steps = make([]int, n)
	steps[0] = 1
	steps[1] = 2
	steps[2] = 4
	for i := 3; i < n; i++ {
		steps[i] = (steps[i-1] + steps[i-2] + steps[i-3]) % 1000000007
	}
	return steps[n-1]
}
