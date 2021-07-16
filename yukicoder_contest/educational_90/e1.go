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

	var p, q, r, k int
	fmt.Fscan(reader, &p, &q, &r, &k)

	a := make([]int, 1000000)
	a[1] = p % 10
	a[2] = q % 10
	a[3] = r % 10

	used := make([]int, 1009)
	used[(a[1]+a[2]+a[3])%10] = 3

	ans := k
	for i := 4; i <= k; i++ {
		a[i] = (a[i-1] + a[i-2] + a[i-3]) % 10
		nex := a[i]*100 + a[i-1]*10 + a[i-2]*1
		if used[nex] != 0 {
			cycle := i - used[nex]
			ans = (k-used[nex])%cycle + used[nex]
			break
		}
		used[nex] = i
	}

	fmt.Fprintf(writer, "%v", a[ans])
}
