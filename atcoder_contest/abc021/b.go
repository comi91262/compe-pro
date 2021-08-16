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
	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var k int
	fmt.Fscan(reader, &k)

	var p = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &p[i])
	}

	m := map[int]int{}
	for i := 0; i < k; i++ {
		m[p[i]]++
	}
	for k, v := range m {
		if v > 1 {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
		if k == a || k == b {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "YES")
}
