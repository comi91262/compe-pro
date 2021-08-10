package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type runLength struct {
	s string
	l int
}

func encodeRunLength(s []string) []runLength {
	r := []runLength{}

	cnt := 0
	for i := 0; i < len(s); i++ {
		if cnt > 0 && r[cnt-1].s == s[i] {
			r[cnt-1].l++
		} else {
			r = append(r, runLength{s: s[i], l: 1})
			cnt++
		}
	}

	return r
}

func main() {

	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	r := encodeRunLength(s)

	for i := range r {
		fmt.Fprintf(writer, "%v", r[i].s+strconv.Itoa(r[i].l))

	}
	fmt.Fprintf(writer, "\n")

}
