package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r, g, b string

	fmt.Scan(&r, &g, &b)

	digits, _ := strconv.Atoi(r + g + b)

	if digits%4 == 0 {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")

	}

}
