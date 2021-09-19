package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func minElement(arg ...int) int {
	min := arg[0]
	me := 0
	for i, x := range arg {
		if min > x {
			min = x
			me = i
		}
	}
	return me
}
func maxElement(arg ...int) int {
	max := arg[0]
	me := 0
	for i, x := range arg {
		if max < x {
			max = x
			me = i
		}
	}
	return me
}

func main() {
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	d := []int{a, b, c}

	mxe := maxElement(d...) + 1
	mne := minElement(d...) + 1
	fmt.Fprintf(writer, "%v\n", string("A"[0]+byte(6/(mxe*mne)-1)))
}
