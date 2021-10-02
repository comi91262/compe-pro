package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toNumber(a string, base int) (num int) {
	for i := range a {
		num *= base
		num += int(a[i] - "0"[0])
	}
	return
}

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

	var s string
	fmt.Fscan(reader, &s)

	n := len(s)
	ans := 0
	for i := 0; i < 1<<n; i++ {
		a := ""
		b := ""
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				a += string(s[j])
			} else {
				b += string(s[j])
			}
		}
		if a == "" || b == "" {
			continue
		}
		if a[0] == "0"[0] || b[0] == "0"[0] {
			continue
		}
		c := strings.Split(a, "")
		sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })
		a = strings.Join(c, "")
		d := strings.Split(b, "")
		sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
		b = strings.Join(d, "")

		an := toNumber(a, 10)
		bn := toNumber(b, 10)

		ans = max(ans, an*bn)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
