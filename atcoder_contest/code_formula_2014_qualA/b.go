package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	p := make([]string, 10)
	for i := 0; i < 10; i++ {
		p[i] = "x"
	}
	for i := 0; i < a; i++ {
		fmt.Scan(&c)
		p[c] = "."
	}
	for i := 0; i < b; i++ {
		fmt.Scan(&c)
		p[c] = "o"
	}
	fmt.Printf("%v %v %v %v\n %v %v %v\n  %v %v\n   %v\n", p[7], p[8], p[9], p[0], p[4], p[5], p[6], p[2], p[3], p[1])
}
