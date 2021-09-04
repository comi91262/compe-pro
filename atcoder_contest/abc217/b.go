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
	var s1 string
	fmt.Fscan(reader, &s1)
	var s2 string
	fmt.Fscan(reader, &s2)
	var s3 string
	fmt.Fscan(reader, &s3)

	m := map[string]int{}
	m[s1]++
	m[s2]++
	m[s3]++

	switch {
	case m["ABC"] == 0:
		fmt.Fprintf(writer, "%v\n", "ABC")
		return
	case m["ARC"] == 0:
		fmt.Fprintf(writer, "%v\n", "ARC")
		return
	case m["AGC"] == 0:
		fmt.Fprintf(writer, "%v\n", "AGC")
		return
	case m["AHC"] == 0:
		fmt.Fprintf(writer, "%v\n", "AHC")
		return
	}

}
