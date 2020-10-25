package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var (
		m   = make(map[int]int, len(arr2))
		ret []int
		r   []int
	)

	for _, v := range arr2 {
		m[v] = 0
	}

	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			r = append(r, v)
		}
	}
	for _, v := range arr2 {
		for i := 0; i < m[v]; i++ {
			ret = append(ret, v)
		}
	}
	sort.Ints(r)

	ret = append(ret, r...)

	return ret
}
