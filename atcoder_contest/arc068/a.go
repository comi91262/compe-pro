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

	ans := x / 11

	switch {
	case x%11 == 0:
		fmt.Fprintf(writer, "%v\n", 2*ans)
	case x%11 <= 6:
		fmt.Fprintf(writer, "%v\n", 2*ans+1)
	default:
		fmt.Fprintf(writer, "%v\n", 2*ans+2)
	}

	//for x := 0; x < 30; x++ {
	//	ans := x / 11
	//	switch {
	//	case x%11 == 0:
	//		fmt.Fprintf(writer, "%v %v\n", x, 2*ans)
	//	case x%11 <= 6:
	//		fmt.Fprintf(writer, "%v %v\n", x, 2*ans+1)
	//	default:
	//		fmt.Fprintf(writer, "%v %v\n", x, 2*ans+2)
	//	}
	//}
	// }
}
