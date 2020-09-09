package main

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	if len(nums) < 1 {
		return ret
	}
	i, j := 0, k-1
	var max int
	for ; i <= len(nums)-k; i++ {
		if i == 0 || nums[i-1] == max {
			max = nums[i]
			for l := i; l <= j && l < len(nums); l++ {
				if nums[l] > max {
					max = nums[l]
				}
			}
		} else {
			if nums[j] > max {
				max = nums[j]
			}
		}

		ret = append(ret, max)
		j++
	}
	return ret
}
