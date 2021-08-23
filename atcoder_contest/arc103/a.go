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
	i int
	n int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
	}

	odd := map[int]int{}
	even := map[int]int{}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even[v[i]]++
		} else {
			odd[v[i]]++
		}
	}
	o := []pair{}
	e := []pair{}

	for k, v := range odd {
		o = append(o, pair{k, v})
	}
	for k, v := range even {
		e = append(e, pair{k, v})
	}
	sort.Slice(o, func(i, j int) bool { return o[i].n > o[j].n })
	sort.Slice(e, func(i, j int) bool { return e[i].n > e[j].n })

	if o[0].i == e[0].i {
		if len(o) == 1 {
			if len(e) == 1 {
				fmt.Fprintf(writer, "%v\n", n/2)
			} else {
				fmt.Fprintf(writer, "%v\n", n/2-e[1].n)
			}
			return
		}

		if len(e) == 1 {
			if len(o) == 1 {
				fmt.Fprintf(writer, "%v\n", n/2)
			} else {
				fmt.Fprintf(writer, "%v\n", n/2-o[1].n)
			}
			return
		}

		if o[0].n > e[0].n {
			fmt.Fprintf(writer, "%v\n", n-o[0].n-e[1].n)
		} else if o[0].n < e[0].n {
			fmt.Fprintf(writer, "%v\n", n-e[0].n-o[1].n)
		} else {
			if o[1].n > e[1].n {
				fmt.Fprintf(writer, "%v\n", n-e[0].n-o[1].n)
			} else {
				fmt.Fprintf(writer, "%v\n", n-o[0].n-e[1].n)
			}
		}
	} else {
		fmt.Fprintf(writer, "%v\n", n-o[0].n-e[0].n)
	}

}
