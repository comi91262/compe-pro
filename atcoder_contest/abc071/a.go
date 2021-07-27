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
	var x int
	fmt.Fscan(reader, &x)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

    if abs(x - a) >= abs(x - b){
        fmt.Fprintf(writer, "%v\n", "B")
    } else {
        fmt.Fprintf(writer, "%v\n", "A")
    }

}

func abs(x int) int {
    if x >= 0 {
        return x
    } else {
        return x * -1
    }
}
