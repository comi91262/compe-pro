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
	n := 6

	me := map[int]int{}
	var e = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &e[i])
		me[e[i]]++
	}
	var b int
	fmt.Fscan(reader, &b)
	ml := map[int]int{}
	var l = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
		ml[l[i]]++
	}

	cnt := 0
	for k := range me {
		if ml[k] > 0 {
			cnt++
			continue
		}
	}
	ok := false
	if me[b] > 0 || ml[b] > 0 {
		ok = true
	}

	switch cnt {
	case 6:
		fmt.Fprintf(writer, "%v\n", 1)
	case 5:
		if ok {
			fmt.Fprintf(writer, "%v\n", 2)
		} else {
			fmt.Fprintf(writer, "%v\n", 3)
		}
	case 4:
		fmt.Fprintf(writer, "%v\n", 4)
	case 3:
		fmt.Fprintf(writer, "%v\n", 5)
	default:
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
