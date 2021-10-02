package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	m := map[int]int{}
	m[a]++
	m[b]++
	m[c]++
	if m[1] > m[2] {
		fmt.Fprintf(writer, "%v\n", 1)
	} else {
		fmt.Fprintf(writer, "%v\n", 2)
	}

}
