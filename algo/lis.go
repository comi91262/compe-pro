package algo

// LIS (Longest increase string: 最長増加文字列) の文字数を求める関数
// 配列のi番目の要素は i番目の数を含む最長増加文字列の文字数を返す
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
// binarySearch(search.go)が必要
func lis(a []int) []int {
	r := make([]int, len(a))

	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := binarySearch(b, a[i], lowerBound)

		if cnt == len(b) {
			b = append(b, a[i])
		} else {
			b[cnt] = a[i]
		}
		r[i] = cnt + 1
	}

	return r
}
