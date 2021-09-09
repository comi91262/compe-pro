package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  int
	second int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i].first, &a[i].second)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].first > a[j].first })

	fmt.Fprintf(writer, "%v\n", a[0].first+a[0].second)

}
