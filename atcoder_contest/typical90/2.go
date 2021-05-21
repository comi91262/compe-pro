package main

import "fmt"

func printPram(n, max, left, right int, s string) {
	if left > max/2 {
		return
	}
	if right > max/2 {
		return
	}
	if left < right {
		return
	}
	if left == max {
		return
	}
	if right == max {
		return
	}

	if n == 0 {
		fmt.Println(s)
		return
	}

	printPram(n-1, max, left+1, right, s+"(")
	printPram(n-1, max, left, right+1, s+")")
}

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 != 0 {
		return
	}

	printPram(n, n, 0, 0, "")
}
