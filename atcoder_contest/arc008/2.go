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
	var M int
	fmt.Fscan(reader, &M)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	fmt.Fscan(reader, &ss)
	var t = strings.Split(ss, "")

	mm := map[string]int{}
	for i := 0; i < len(s); i++ {
		mm[s[i]]++
	}
	m := map[string]int{}
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	ans := 0
	for k, v := range mm {
		if m[k] == 0 {
			fmt.Fprintf(writer, "%v\n", -1)
			return
		}
		ans = max(ans, (v+m[k]-1)/m[k])
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
