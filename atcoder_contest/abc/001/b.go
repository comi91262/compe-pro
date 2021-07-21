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

	var m int
	fmt.Fscan(reader, &m)

	switch {
	case m < 100:
		fmt.Fprintf(writer, "00\n")
	case 100 <= m && m < 1000:
		fmt.Fprintf(writer, "0%v\n", m/100)
	case 1000 <= m && m <= 5000:
		fmt.Fprintf(writer, "%v\n", m/100)
	case 6000 <= m && m <= 30000:
		fmt.Fprintf(writer, "%v\n", m/1000+50)
	case 35000 <= m && m <= 70000:
		fmt.Fprintf(writer, "%v\n", (m/1000-30)/5+80)
	case 70000 < m:
		fmt.Fprintf(writer, "%v\n", 89)
	}
}
