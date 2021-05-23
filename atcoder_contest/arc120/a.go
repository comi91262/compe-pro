package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [200000]int
var b [200000]int
var c [200000]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			b[i] = a[i]
			c[i] = a[i]
		} else {
			b[i] = a[i] + b[i-1]
			//	if a[i-1]+b[i-1] >= b[i] {
			//		c[i] = c[i-1]
			//	} else {
			//	}
			if c[i-1] > a[i] {
				c[i] = c[i-1]
			} else {
				c[i] = a[i]
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += b[i]
		fmt.Fprintln(writer, sum+c[i]*(i+1))
	}

}

