package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	s string
	n int
}

type Deque struct {
	cl   int
	cr   int
	data []pair
}

func (d *Deque) pushFront(x pair) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x pair) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() pair {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() pair {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) size() int {
	return d.cr - d.cl
}

func (d *Deque) empty() bool {
	return d.size() == 0
}

func (d *Deque) get(x int) pair {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]pair, size*2, size*2)}
}

func main() {
	defer writer.Flush()
	var q int
	fmt.Fscan(reader, &q)

	d := create(100001)

	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		switch t {
		case 1:
			var c string
			fmt.Fscan(reader, &c)
			var x int
			fmt.Fscan(reader, &x)
			d.pushBack(pair{c, x})
		case 2:
			var D int
			fmt.Fscan(reader, &D)
			m := map[string]int{}
			for D > 0 && d.size() > 0 {
				tmp := d.popFront()
				if tmp.n > D {
					m[tmp.s] += D
					d.pushFront(pair{tmp.s, tmp.n - D})
					D = 0
				} else {
					m[tmp.s] += tmp.n
					D -= tmp.n
				}
			}
			cnt := 0
			for i := 0; i < 26; i++ {
				cnt += m[string("a"[0]+byte(i))] * m[string("a"[0]+byte(i))]
			}
			fmt.Fprintf(writer, "%v\n", cnt)
		}
	}
}

