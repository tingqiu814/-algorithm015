package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret = [][]int{}
	var set = map[string]bool{}
	for i, v := range nums {
		target := v * -1

		var j, k = i + 1, len(nums) - 1
		for j < k {
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				if _, ok := set[fmt.Sprintf("%v-%v-%v", nums[i], nums[j], nums[k])]; !ok {
					set[fmt.Sprintf("%v-%v-%v", nums[i], nums[j], nums[k])] = true
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
				j++
			}
		}

	}
	return ret
}
