package main

func maxArea(height []int) int {
	var ret = 0
	//var getArea func (start, end int) int
	var getArea = func(start, end int) int {
		hei := 0
		if height[start] > height[end] {
			hei = height[end]
		} else {
			hei = height[start]
		}
		var weight = end - start
		return hei * weight

	}
	for i, j := 0, len(height)-1; i < j; {
		if getArea(i, j) > ret {
			ret = getArea(i, j)
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}
