package main

import "strconv"

func numDecodings(s string) int {
	var cnt = 0
	if len(s) == 1 {
		if s != "0" {
			cnt += 1
		}
	}
	if len(s) >= 2 {
		n, _ := strconv.Atoi(s[0:2])
		if n <= 26 {
			cnt += numDecodings(s[1:]) + numDecodings(s[2:])
		}
	}
	return cnt
}
