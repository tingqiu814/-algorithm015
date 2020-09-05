package main

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ret = []int{}
	var index1, index2 = 0, 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			ret = append(ret, nums1[index1])
			index1++
			index2++
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			index2++
		}
	}
	return ret
}
