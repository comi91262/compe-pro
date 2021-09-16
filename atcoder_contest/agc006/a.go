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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	mx := 0
	for i := 0; i < len(s); i++ {
		cnt := 0
		for j := 0; j < len(t); j++ {
			if i+j < n && s[i+j] == t[j] {
				cnt++
			} else {
				break
			}
		}
		mx = max(mx, cnt)
	}

	fmt.Fprintf(writer, "%v\n", 2*n-mx)
}
