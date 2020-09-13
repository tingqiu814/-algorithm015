package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	//comb := []int{}
	//var dfs func(target, idx int)
	//dfs = func(target, idx int) {
	//	if idx == len(candidates) {
	//		return
	//	}
	//	if target == 0 {
	//		ans = append(ans, append([]int(nil), comb...))
	//		return
	//	}
	//	// 直接跳过
	//	dfs(target, idx+1)
	//	// 选择当前数
	//	if target-candidates[idx] >= 0 {
	//		comb = append(comb, candidates[idx])
	//		dfs(target-candidates[idx], idx)
	//		comb = comb[:len(comb)-1]
	//	}
	//}
	//dfs(target, 0)
	//默写
	resultTmp := []int{}
	var dfs2 func(target, idx int)
	dfs2 = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			// 需要append方式重新生成一个数组
			ans = append(ans, append([]int(nil), resultTmp...))
			return
		}
		if target-candidates[idx] >= 0 {
			resultTmp = append(resultTmp, candidates[idx])
			dfs2(target-candidates[idx], idx)
			resultTmp = resultTmp[:len(resultTmp)-1]
		}
		// 跳过当前值向下一位
		dfs2(target, idx+1)
	}
	dfs2(target, 0)
	return ans
}
