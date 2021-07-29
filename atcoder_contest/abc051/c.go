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

	var sx int
	fmt.Fscan(reader, &sx)
	var sy int
	fmt.Fscan(reader, &sy)
	var tx int
	fmt.Fscan(reader, &tx)
	var ty int
	fmt.Fscan(reader, &ty)

	dx := tx - sx
	dy := ty - sy

	for i := 0; i < dy; i++ {
		fmt.Fprintf(writer, "U")
	}
	for i := 0; i < dx; i++ {
		fmt.Fprintf(writer, "R")
	}

	for i := 0; i < dy; i++ {
		fmt.Fprintf(writer, "D")
	}
	for i := 0; i < dx; i++ {
		fmt.Fprintf(writer, "L")
	}

	fmt.Fprintf(writer, "L")
	for i := 0; i < dy+1; i++ {
		fmt.Fprintf(writer, "U")
	}
	for i := 0; i < dx+1; i++ {
		fmt.Fprintf(writer, "R")
	}
	fmt.Fprintf(writer, "D")
	fmt.Fprintf(writer, "R")
	for i := 0; i < dy+1; i++ {
		fmt.Fprintf(writer, "D")
	}
	for i := 0; i < dx+1; i++ {
		fmt.Fprintf(writer, "L")
	}
	fmt.Fprintf(writer, "U")

	fmt.Fprintf(writer, "\n")
}
