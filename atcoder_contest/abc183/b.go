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

	var sx float64
	fmt.Fscan(reader, &sx)
	var sy float64
	fmt.Fscan(reader, &sy)
	var gx float64
	fmt.Fscan(reader, &gx)
	var gy float64
	fmt.Fscan(reader, &gy)

	gy = -gy
	dx := gx - sx
	dy := gy - sy

    fmt.Fprintf(writer, "%v\n", sx+sy*dx/dy*-1.0)
}
