package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n float64
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	m := map[int]int{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "A":
			m[4]++
		case "B":
			m[3]++
		case "C":
			m[2]++
		case "D":
			m[1]++
		case "F":
		}
	}
	fmt.Fprintf(writer, "%v\n", float64(m[4]*4+m[3]*3+m[2]*2+m[1]*1)/n)
}

