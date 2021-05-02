package main

import (
	"fmt"
)

// func range(input string){
//     for _, c := range input {
//         fmt.Printf("%c ", c)
//         fmt.Println(reflect.TypeOf(c))
//     }
//     fmt.Println
// }

func divide_n(n int) int {
	result := 0
	for {
		if n%2 != 0 {
			break
		}
		n = n / 2
		result++
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	min := 500000001

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])

		b := divide_n(a[i])
		if b < min {
			min = b
		}
	}

	fmt.Println(min)
}