package main

func firstUniqChar(s string) int {
	t := make(map[int32]int)
	for _, v := range s {
		if _, ok := t[v]; ok {
			t[v] = t[v] + 1
		} else {
			t[v] = 1
		}
	}

	for k, v := range s {
		if i, ok := t[v]; ok && i == 1 {
			return k
		}
	}
	return -1

}
