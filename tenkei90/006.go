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
	var n, k int
	var s string
	fmt.Fscan(reader, &n, &k, &s)

	ss := strings.Split(s, "")
	var c [26][100001]int

	for i := n - 1; i >= 0; i-- {
		ru := ss[i][0]
		for j := 0; j < 26; j++ {
			if j == int(ru-'a') {
				c[j][i] = i + 1
			} else {
				c[j][i] = c[j][i+1]
			}
		}
	}

	ans := ""
	rest := k
	for i := 0; i < n; {
		for j := 0; j < 26; j++ {
			next := c[j][i]
			if next == 0 || n-next < rest-1 {
				continue
			}
			rest--
			i = next
			ans += string(j + 'a')
			break
		}

		if rest == 0 {
			break
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
