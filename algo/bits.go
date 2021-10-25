package algo

func intersection(a, b int) int {
	return a & b
}
func unionSet(a, b int) int {
	return a | b
}
func diff(a, b int) int {
	return a ^ b
}

// 要素xを削除する
func subtract(a, n, x int) int {
	b := 1 << (n - 1 - x)
	return a &^ b
}

func increment(a, n int) int {
	if a&1<<0 != 0 {
		return a>>1 | 1<<(n-1)
	} else {
		return a >> 1
	}
}
func decrement(a, n int) int {
	if a&1<<(n-1) != 0 {
		return a<<1&(1<<n-1) | 1<<0
	} else {
		return a << 1 & (1<<n - 1)
	}
}
