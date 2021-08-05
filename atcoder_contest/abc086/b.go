package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)

	c := a + b
	i, _ := strconv.Atoi(c)

	n := int(math.Sqrt(float64(i)))
	if n*n == i {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
