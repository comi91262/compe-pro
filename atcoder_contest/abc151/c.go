package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type data struct {
	pre string
	p   int
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var p = make([]int, m)
	var s = make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &s[i])
	}

	an := 0
	pn := 0

	var d = make([]data, n+1)
	for i := 0; i < m; i++ {
		if s[i] == "WA" && d[p[i]].pre != "AC" {
			d[p[i]].p++
			d[p[i]].pre = s[i]
		}

		if s[i] == "AC" && d[p[i]].pre != "AC" {
			an++
			pn += d[p[i]].p
			d[p[i]].pre = s[i]
		}
	}

	fmt.Fprintf(writer, "%v %v\n", an, pn)
}
