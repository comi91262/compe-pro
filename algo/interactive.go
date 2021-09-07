package algo

import (
	"fmt"
)

func ask() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("? 1 %v\n", i)
		var c int
		fmt.Scan(&c)
	}

	fmt.Printf("! %v\n", 1)
}
