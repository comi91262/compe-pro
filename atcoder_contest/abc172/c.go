package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) pushFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() int {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() int {
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

func (d *Deque) get(x int) int {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]int, size*2, size*2)}
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var k int
	fmt.Fscan(reader, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	d := create(n + m)

	tot := 0
	for d.size() < n && tot+a[d.size()] <= k {
		tot += a[d.size()]
		d.pushBack(a[d.size()])
	}

	ans := d.size()
	an := d.size()
	for i := 0; i < m; i++ {
		d.pushFront(b[i])
		tot += b[i]

		for !d.empty() && tot > k {
			tot -= d.popBack()
			an--
		}
		if an < 0 {
			break
		}

		ans = max(ans, d.size())
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
