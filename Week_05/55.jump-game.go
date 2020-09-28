package main

func canJump(nums []int) bool {
	var post = len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= post {
			post = i
		}
	}
	return post <= 0
}
