package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func mex(arg ...int) int {
	m := map[int]int{}
	for _, x := range arg {
		m[x]++
	}

	for i := 0; ; i++ {
		if m[i] == 0 {
			return i
		}
	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := map[int]int{}
	idx := 0
	for i := 0; i < n; i++ {
		m[a[i]]++
		for i := idx; ; i++ {
			if m[i] == 0 {
				idx = i
				break
			}
		}
		fmt.Fprintf(writer, "%v\n", idx)
	}

}
