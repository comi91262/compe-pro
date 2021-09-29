package algo

// うるう年を計算する関数
func isLeapYear(y int) (isOk bool) {
	isOk = y%4 == 0
	isOk = isOk && y%100 != 0
	isOk = isOk || y%400 == 0
	return
}

// 周期性を求める問題を解く関数
func solveCycle(n, st, mod int) {
	used := make([]int, mod)
	used[st] = 1
	start := 0
	pre := 0

	// 周期の開始から一周後の開始地点を求める
	for i := 2; i <= n; i++ {
		nex := pre // ...
		if used[nex] != 0 {
			start = i
			break
		}
		used[nex] = i
		pre = nex
	}

	n -= start
	if start == 0 {
		return
	}

	cycle := []int{}
	used = make([]int, mod)
	// 周期を求める
	for i := 1; i <= n; i++ {
		nex := pre
		if used[nex] != 0 {
			break
		}
		cycle = append(cycle, nex)
		used[nex] = i
		pre = nex
	}

}
