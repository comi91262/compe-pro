package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	cnt := 0
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", n-1-max(0, cnt-2*k))

}

