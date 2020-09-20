package main

func search2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high-1 {
		mid := (high + low) / 2
		a, b, c := nums[low], nums[mid], nums[high]
		if b == target {
			return mid
		}
		switch {
		// 上面两个case在递增序列找同二分查找
		case target >= a && target <= b:
			high = mid
		case target >= b && target <= c:
			low = mid
			// 如果在递增的区间没有找到
			// 那么只可能在递减的区间中
			// 由于是旋转数组，只可能有一个递减区间
		case a > b:
			high = mid
		case b > c:
			low = mid
		default:
			return -1
		}
	}
	if nums[low] == target {
		return low
	}
	if nums[high] == target {
		return high
	}
	return -1
}

func search(nums []int, target int) int {
	// 由于是半单调有序的，并且有边界，可以用二分查找
	var (
		left  = 0
		right = len(nums) - 1
	)
	if len(nums) == 0 {
		return -1
	}
	for left < right {
		var mid = (right + left) / 2
		if target == nums[mid] {
			return mid
		}
		switch {
		// 上面两个case在递增序列找同二分查找
		case target >= nums[left] && target <= nums[mid]:
			right = mid
		case target >= nums[mid] && target <= nums[right]:
			left = mid
			// 如果在递增的区间没有找到
			// 那么只可能在递减的区间中
			// 由于是旋转数组，只可能有一个递减区间
		case nums[left] > nums[mid]:
			right = mid
		case nums[mid] > nums[right]:
			left = mid
		default:
			return -1
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}

	return -1
}
