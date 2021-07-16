package algo

import "errors"

// n進数変換に関する処理

// '0' -> 0
func byteToDigit(c byte) (int, error) {
	n := c - 48

	if n < 0 || 9 < n {
		return int(c), errors.New("wrong number")
	}

	return int(n), nil
}

// 123 -> [1,2,3]
func strToDigits(s string) []int {
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i], _ = byteToDigit(s[i])
	}
	return r
}

// base進数で xを桁ごとに分解する (1から10進数まで)
// 10(10) -> [1, 0]
// 10(2)  -> [1, 0, 1, 0]
func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)
	for x != 0 {
		r = append(r, x%base)
		x = x / base
	}

	reverse(r, 0, len(r)-1)

	return r
}

// base進数でaである数字列を10進数の数字に直す, toDigitsの逆
//
//  [1,0,1,0], 2 -> 10
//  [1,0], 10 -> 10
func toNumber(a []int, base int) int {
	cnt := 1
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[len(a)-1-i] * cnt
		cnt *= base
	}
	return ans
}
