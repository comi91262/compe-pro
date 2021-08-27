package algo

func intersect(a, b map[string]int) map[string]int {
	r := map[string]int{}
	for k, v := range a {
		if b[k] > 0 {
			r[k] = min(b[k], v)
		}
	}

	return r
}

func union(a, b map[int]int) map[int]int {
	r := map[int]int{}
	for k, v := range a {
		r[k] += v
	}
	for k, v := range b {
		r[k] += v
	}

	return r
}
