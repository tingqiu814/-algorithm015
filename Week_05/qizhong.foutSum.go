package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ret = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			var subTarget = target - nums[i] - nums[j]
			left := j + 1
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] == subTarget {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					break
				} else if nums[left]+nums[right] > subTarget {
					right -= 1
				} else {
					left += 1
				}

			}
		}
	}
	return ret
}
