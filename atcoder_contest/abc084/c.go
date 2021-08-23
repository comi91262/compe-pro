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

	var c = make([]int, n-1)
	var s = make([]int, n-1)
	var f = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &c[i])
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &f[i])
	}

	for j := 0; j < n; j++ {
		ans := 0
		for i := j; i < n-1; i++ {
			if i == j {
				ans = s[i] + c[i]
				continue
			}

			if ans < s[i] {
				ans = s[i]
			} else {
				if (ans-s[i])%f[i] != 0 {
					ans += f[i] - (ans-s[i])%f[i]
				}
			}
			ans += c[i]

		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
