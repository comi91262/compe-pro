package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var r float64
	fmt.Fscan(reader, &r)
	var d float64
	fmt.Fscan(reader, &d)

	fmt.Fprintf(writer, "%v\n", 2.0*math.Pi*d*math.Pi*r*r)

}
