package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var p int
	fmt.Fscan(reader, &p)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]%2]++
	}
	if p == 1 && m[1] == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	c1 := 0
	c2 := 0
	if p == 0 {
		if m[0] > 0 {
			c1 += pow(2, m[0])
		}
		if m[1] > 0 {
			c2 += pow(2, m[1]-1)
		}
	} else {
		if m[1] > 0 {
			c1 += pow(2, m[1])
		}
		if m[0] > 0 {
			c2 += pow(2, m[0]-1)
		}
	}
	if c1 == 0 || c2 == 0 {
		fmt.Fprintf(writer, "%v\n", c1+c2)
	} else {
		fmt.Fprintf(writer, "%v\n", c1*c2)
	}
}
