package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)
	r := regexp.MustCompile(`^([a-z]+keyence$|k[a-z]+eyence$|ke[a-z]+yence$|key[a-z]+ence$|keye[a-z]+nce$|keyen[a-z]+ce$|keyenc[a-z]+e$|keyence[a-z]+$|keyence$)`)

	if r.MatchString(s) {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}
}
