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

	var n string
	fmt.Fscan(reader, &n)

	r := []rune(n)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	n = string(r)

	o, e := 0, 0
	for i := 0; i < len(n); i++ {
		if i%2 == 0 {
			o += int(n[i] - "0"[0])
		} else {
			e += int(n[i] - "0"[0])
		}
	}
	fmt.Fprintf(writer, "%v %v\n", e, o)
}
