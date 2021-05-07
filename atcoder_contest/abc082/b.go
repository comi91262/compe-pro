package main

import (
	"fmt"
	"sort"
)

func cond1(x, y string) bool {
	n := len(x)
	m := len(y)

	if n < m {
		for i := 0; i < n; i++ {
			if x[i] != y[i] {
				return false
			}
		}
		return true
	}

	return false
}

func cond2(x, y string) bool {
	n := len(x)
	m := len(y)

	var l int
	if n > m {
		l = m
	} else {
		l = n
	}

	for i := 0; i < l; i++ {
		if x[i] == y[i] {
			continue
		}

		return x[i] < y[i]
	}

	return false
}

func main() {
	var s, t string

	fmt.Scan(&s, &t)

	sRune := []rune(s)
	tRune := []rune(t)

	sort.Slice(sRune, func(i, j int) bool { return sRune[i] < sRune[j] })
	sort.Slice(tRune, func(i, j int) bool { return tRune[i] > tRune[j] })

	s = string(sRune)
	t = string(tRune)

	if cond1(s, t) || cond2(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
