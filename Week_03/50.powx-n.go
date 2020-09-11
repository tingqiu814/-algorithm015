package main

func myPow(x float64, n int) float64 {
	var ret float64
	if n < 0 {
		n = n * -1
		x = 1 / x
	}
	if n == 0 {
		return 1
	}
	var halfPow = myPow(x, n/2)
	if n%2 == 0 {
		ret = halfPow * halfPow
	} else {
		ret = halfPow * halfPow * x
	}
	return ret
}
