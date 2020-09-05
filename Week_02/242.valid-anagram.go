package main

func isAnagram(s string, t string) bool {
	var sMap = map[rune]int{}
	for _, item := range s {
		sMap[item]++
	}
	for _, item := range t {
		sMap[item]--
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}
