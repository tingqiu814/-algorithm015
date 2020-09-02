package main

// 优化到100% 要用到一些go的特性，如果字符串+=的话会不停复制占用内存，所以用slice
func removeOuterParentheses(S string) string {
	var stackNum = 0
	var ret = []uint8{}
	for i := 0; i < len(S); i++ {
		if S[i] == uint8(41) {
			stackNum--
			if stackNum != 0 {
				ret = append(ret, S[i])
			}
		} else {
			if stackNum != 0 {
				ret = append(ret, S[i])
			}
			stackNum++
		}
	}
	return string(ret)
}
