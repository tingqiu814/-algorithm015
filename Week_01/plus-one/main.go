package main

func main() {
	var input = []int{1, 2, 3}
	var ret = plusOne(input)
	print(ret)
}
func plusOne(digits []int) []int {
	var carry = false
	for i := len(digits) - 1; i >= 0; i-- {
		if carry || i == len(digits)-1 {
			if digits[i]+1 == 10 {
				digits[i] = 0
				carry = true
			} else {
				digits[i] = digits[i] + 1
				carry = false
			}
		}
	}
	if carry {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}
