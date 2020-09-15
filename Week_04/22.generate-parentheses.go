package main

func generateParenthesis(n int) []string {
	ret := []string{}
	var _gen func(left, right int, s string)
	_gen = func(left, right int, s string) {
		if left == right && left == n {
			ret = append(ret, s)
		}
		if left < n {
			_gen(left+1, right, s+"(")
		}

		if left > right {
			_gen(left, right+1, s+")")
		}
	}
	_gen(0, 0, "")
	return ret
}

func _generate(left int, right int, max int, s string, Output *[]string) {
	// 递归终止条件
	if left == right && left == max {
		*Output = append(*Output, s)
		return
	}
	// 递归主体
	if left < max {
		_generate(left+1, right, max, s+"(", Output)
	}
	if right < left {
		_generate(left, right+1, max, s+")", Output)
	}
}
