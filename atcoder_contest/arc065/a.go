package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := ""
	for {
		if strings.HasSuffix(s, "dreamer"+t) {
			t = "dreamer" + t

		} else if strings.HasSuffix(s, "dream"+t) {
			t = "dream" + t

		} else if strings.HasSuffix(s, "eraser"+t) {
			t = "eraser" + t

		} else if strings.HasSuffix(s, "erase"+t) {
			t = "erase" + t
		} else {
			if s == t {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			break
		}
	}

}