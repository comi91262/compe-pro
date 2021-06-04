package main

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func combination(n int, r int) int {
	return fact(n) / fact(r) * fact(n-r)
}
