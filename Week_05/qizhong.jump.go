package main

func jump(nums []int) int {
	var ret = 0
	position := len(nums) - 1
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				ret++
			}

		}
	}

	return ret
}
