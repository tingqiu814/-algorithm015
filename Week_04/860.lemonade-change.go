package main

func lemonadeChange(bills []int) bool {
	var own = map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	var check func(get int) bool
	check = func(get int) bool {
		own[get] += 1
		get = get - 5
		for get >= 20 && own[20] > 0 {
			own[20] -= 1
			get -= 20
		}
		for get >= 10 && own[10] > 0 {
			own[10] -= 1
			get -= 10
		}
		for get >= 5 && own[5] > 0 {
			own[5] -= 1
			get -= 5
		}

		return get == 0
	}
	for i := 0; i < len(bills); i++ {
		b := check(bills[i])
		if !b {
			return false
		}
	}
	return true
}
