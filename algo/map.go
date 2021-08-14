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
