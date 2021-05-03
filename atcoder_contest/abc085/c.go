package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	// a := make([]int, n)
	rvalue := 10000*n - y
	for j := 0; j <= n; j++ {
		for k := 0; k <= n; k++ {

			if 5000*j+9000*k == rvalue {
				if n-j-k < 0 {
					break
				}
				fmt.Printf("%d %d %d\n", n-j-k, j, k)
				return

			}

		}
	}

	fmt.Print("-1 -1 -1\n")

}