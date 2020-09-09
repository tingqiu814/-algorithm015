package main

func replaceSpace(s string) string {
	var ret = ""
	for _, v := range s {
		if string(v) != " " {
			ret += string(v)
		} else {
			ret += "%20"
		}
	}

	return ret
}
