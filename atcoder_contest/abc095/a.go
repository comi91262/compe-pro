package main

import (
	"fmt"
)

func main() {
	var orders string

	fmt.Scan(&orders)

	total := 700
	for _, order := range orders {
		if order == 'o' {
			total += 100
		}
	}

	fmt.Println(total)
}
