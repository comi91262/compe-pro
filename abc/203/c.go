package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]struct{ n, y int }, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].n, &a[i].y)
	}

	sort.Slice(a, func(i, j int) bool { return a[i].n < a[j].n })

	sum := k
	for i := 0; i < n; i++ {
		if sum < a[i].n {
			break
		}
		sum += a[i].y
	}

	fmt.Fprintf(writer, "%d\n", sum)
}
