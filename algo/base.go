package algo

// 123 -> [3,2,1] (10)
func strToDigits(s string) []int {
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		var n = int(s[i] - 48)
		ans[len(s)-1-i] = n
	}
	return ans
}

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	ans := make([]int, 0)
	for x != 0 {
		ans = append(ans, x%base)
		x = x / base
	}
	return ans
}

func toNumber(a []int, base int) int {
	cnt := 1
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[i] * cnt
		cnt *= base
	}
	return ans
}
