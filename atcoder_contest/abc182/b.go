package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	ans := 0
	mx := 0
	for k := 2; k <= 1000; k++ {
		cnt := 0
		for i := 0; i < n; i++ {
			if a[i]%k == 0 {
				cnt++
			}
		}
		if cnt > mx {
			mx = cnt
			ans = k
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

