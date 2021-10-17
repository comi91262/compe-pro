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

func digit(x int) int {
	r := 0
	for x > 0 {
		r++
		x /= 10
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

	return r
}

// base進数で xを桁ごとに分解する 負の進数にも対応
func toMathDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)

	for x != 0 {
		r = append(r, mod(x, base))
		if x < 0 {
			x = (x - 1) / base
		} else {
			x = x / base
		}
	}

	reverse(r, 0, len(r)-1)

	return r
}

// base進数でaである文字列を10進数の数字に直す
//
//  "1010", 2 -> 10
//  "10", 10 -> 10
func toNumber(a string, base int) (num int) {
	for i := range a {
		num *= base
		num += int(a[i] - '0')
	}
	return
}

// 2進数でxを表し, 各桁の数をスライスで返す
// size の分0詰めする
func toBinaryDigits(x, size int) []int {
	r := make([]int, size)
	for i := 0; i < size; i++ {
		if x == 0 {
			break
		}
		r[size-1-i] = x % 2
		x /= 2
	}
	return r
}
