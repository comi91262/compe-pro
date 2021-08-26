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

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b1 string
	fmt.Fscan(reader, &b1)
	b2 := strings.Split(b1, ".")
	b3 := make([]int, 2)
	b3[0], _ = strconv.Atoi(b2[0])
	b3[1], _ = strconv.Atoi(b2[1])

	fmt.Fprintf(writer, "%v\n", a*(b3[0]*100+b3[1])/100)

}
