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

	var n int
	fmt.Fscan(reader, &n)

	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	var t = make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &t[i])
	}

	ma1 := map[string]int{}
	for i := range s {
		ma1[s[i]]++
	}

	ma2 := map[string]int{}
	for i := range t {
		ma2[t[i]]++
	}

	ans := 0
	for k, v := range ma1 {
		ans = max(ans, v-ma2[k])
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
