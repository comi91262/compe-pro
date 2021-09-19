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
	var n int
	fmt.Fscan(reader, &n)
	a := []int{}
	for n > 0 {
		a = append(a, n%36)
		n /= 36
	}
	if len(a) == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}

	for i := 0; i < len(a); i++ {
		switch {
		case 0 <= a[i] && a[i] <= 9:
			fmt.Fprintf(writer, "%v", a[i])
		default:
			fmt.Fprintf(writer, "%v", string(a[i]-10+int("A"[0])))
		}
	}
	fmt.Fprintf(writer, "\n")
}
