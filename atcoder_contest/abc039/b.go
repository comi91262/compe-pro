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

	var xxxx int
	fmt.Fscan(reader, &xxxx)

	xx := int(math.Sqrt(float64(xxxx)))
	x := int(math.Sqrt(float64(xx)))

	fmt.Fprintf(writer, "%v\n", x)
}
