package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}

	for a1 := -1; a1 < 100; a1++ {
		for a2 := -1; a2 < 100; a2++ {
			for a3 := -1; a3 < 100; a3++ {

				var a = []int{a1, a2, a3}
				var b = []int{c[0][0] - a1, c[0][1] - a1, c[0][2] - a1}

				ok := true
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						if a[i]+b[j] != c[i][j] {
							ok = false
						}
					}
				}

				if ok {
					//fmt.Fprintf(writer, "%v %v %v\n", a1, a2, a3)
					fmt.Fprintf(writer, "%v\n", "Yes")
					return
				}
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", "No")

}
