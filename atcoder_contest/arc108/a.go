package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func divisor(n int) []int {
	var r []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i*i != n {
				// r = append(r, n/i)
			}
		}
	}
	return r
}

func main() {
	defer writer.Flush()
	var s int
	fmt.Fscan(reader, &s)
	var p int
	fmt.Fscan(reader, &p)

	d := divisor(p)
	for i := range d {
		a, b := d[i], p/d[i]
		if a+b == s {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")

}
