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
	var s string
	fmt.Fscan(reader, &s)

	a := s[0] - 48
	b := s[1] - 48
	c := s[2] - 48
	d := s[3] - 48

	if a+b+c+d == 7 {
		fmt.Fprintf(writer, "%v+%v+%v+%v=7\n", a, b, c, d)
		return
	}
	if a+b+c-d == 7 {
		fmt.Fprintf(writer, "%v+%v+%v-%v=7\n", a, b, c, d)
		return
	}
	if a+b-c+d == 7 {
		fmt.Fprintf(writer, "%v+%v-%v+%v=7\n", a, b, c, d)
		return
	}
	if a+b-c-d == 7 {
		fmt.Fprintf(writer, "%v+%v-%v-%v=7\n", a, b, c, d)
		return
	}
	if a-b-c-d == 7 {
		fmt.Fprintf(writer, "%v-%v-%v-%v=7\n", a, b, c, d)
		return
	}
	if a-b-c+d == 7 {
		fmt.Fprintf(writer, "%v-%v-%v+%v=7\n", a, b, c, d)
		return
	}
	if a-b+c-d == 7 {
		fmt.Fprintf(writer, "%v-%v+%v-%v=7\n", a, b, c, d)
		return
	}
	if a-b+c+d == 7 {
		fmt.Fprintf(writer, "%v-%v+%v+%v=7\n", a, b, c, d)
		return
	}

}
