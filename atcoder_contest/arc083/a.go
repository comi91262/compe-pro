package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)
	var e int
	fmt.Fscan(reader, &e)
	var f int
	fmt.Fscan(reader, &f)

	ma := f / (100 * a)
	mb := f / (100 * b)
	mc := f / c
	md := f / d

	mx := 0.0
	mt := 0
	ms := 0

	for i := 0; i <= ma; i++ {
		for j := 0; j <= mb; j++ {
			for k := 0; k <= mc; k++ {
				for l := 0; l <= md; l++ {
					wa := 100*a*i + 100*b*j
					su := c*k + d*l
					if wa+su > f {
						continue
					}

					if float64(wa)/100.0*float64(e) < float64(su) {
						continue
					}

					conc := float64(100*su) / float64(wa+su)
					if mx <= conc {
						mx = conc
						mt = wa + su
						ms = su
					}
				}
			}
		}
	}
	fmt.Fprintf(writer, "%v %v\n", mt, ms)

}
