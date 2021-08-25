package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	m := map[string]int{}
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		m[tmp]++
	}

	mx := 0
	for _, v := range m {
		mx = max(mx, v)
	}

	a := []string{}
	for k, v := range m {
		if v == mx {
			a = append(a, k)
		}
	}

	sort.Strings(a)
	for _, name := range a {
		fmt.Fprintf(writer, "%v\n", name)

	}

}
