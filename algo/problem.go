package algo

import "sort"

// うるう年を計算する関数
func isLeapYear(y int) (isOk bool) {
	isOk = y%4 == 0
	isOk = isOk && y%100 != 0
	isOk = isOk || y%400 == 0
	return
}

// 順位づけをする関数 (nlogn)
// input: [100, 70, 90, 80, 90]
// output: map{100:1,90:2,80:4,70:1}
func makeRank(a []int) map[int]int {
	m := map[int]int{}
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	sum := 0
	for i := 0; i < len(m); i++ {
		cnt := m[keys[i]]
		m[keys[i]] = sum + 1
		sum += cnt
	}
	return m
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
