package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)

	t := "2018" + s[4:10]

	fmt.Println(t)
}