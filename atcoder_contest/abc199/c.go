package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	var q int
	fmt.Fscan(reader, &q)

	var t = make([]int, q)
	var a = make([]int, q)
	var b = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &t[i])
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		a[i]--
		b[i]--
	}

	cnt := 0
	pre := 0
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			if pre%2 == 1 {
				var x int
				if a[i] < n {
					x = a[i] + n
				} else {
					x = a[i] - n
				}
				var y int
				if b[i] < n {
					y = b[i] + n
				} else {
					y = b[i] - n
				}
				s[x], s[y] = s[y], s[x]
			} else {
				s[a[i]], s[b[i]] = s[b[i]], s[a[i]]
			}
		case 2:
			cnt++
			pre++
		}
	}

	if cnt%2 == 0 {
		fmt.Fprintf(writer, "%v\n", strings.Join(s, ""))
	} else {
		ans := strings.Join(s, "")
		fmt.Fprintf(writer, "%v\n", ans[n:]+ans[:n])
	}

}
