package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var a [800][800]int
var min = 1000000001

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	mid := k*k/2 + 1
	for i := 0; i < n-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			b := make([]int, k*k)
			cnt := 0
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					b[cnt] = a[x][y]
					cnt++
				}
			}
			sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
			if min > b[mid-1] {
				min = b[mid-1]
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", min)
}
