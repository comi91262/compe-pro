package algo

import "sort"

// LIS (Longest increase subsequence: 最長増加部分列) の文字数を求める関数
// 配列のi番目の要素は i番目の数を含む最長増加部分列の文字数を返す
// 例
// [3 1 4 1 5 9 2 6]
//  3                -> 1
//  1                -> 1
//  1 4              -> 2
//  1                -> 1
//  1 4   5          -> 3
//  1 4   5 9        -> 4
//  1 2              -> 2
//  1 2   5     6    -> 4
func lis(a []int) (lis []int) {
	lis = make([]int, len(a))

	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := sort.Search(len(b), func(j int) bool { return a[i] <= b[j] })

		if cnt == len(b) {
			b = append(b, a[i])
		} else {
			b[cnt] = a[i]
		}
		lis[i] = cnt + 1
	}

	return
}
